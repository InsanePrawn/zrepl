package zfs

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
	"time"
)

func fsvlist(fsv ...string) (r []FilesystemVersion) {

	r = make([]FilesystemVersion, len(fsv))
	for i, f := range fsv {

		// parse the id from fsvlist. it is used to derivce Guid,CreateTXG and Creation attrs
		split := strings.Split(f, ",")
		if len(split) != 2 {
			panic("invalid fsv spec")
		}
		id, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}

		if strings.HasPrefix(f, "#") {
			r[i] = FilesystemVersion{
				Name:      strings.TrimPrefix(f, "#"),
				Type:      Bookmark,
				Guid:      uint64(id),
				CreateTXG: uint64(id),
				Creation:  time.Unix(0, 0).Add(time.Duration(id) * time.Second),
			}
		} else if strings.HasPrefix(f, "@") {
			r[i] = FilesystemVersion{
				Name:      strings.TrimPrefix(f, "@"),
				Type:      Snapshot,
				Guid:      uint64(id),
				CreateTXG: uint64(id),
				Creation:  time.Unix(0, 0).Add(time.Duration(id) * time.Second),
			}
		} else {
			panic("invalid character")
		}
	}
	return
}

func doTest(t *testing.T, left, right []FilesystemVersion, validate func(d FilesystemDiff)) {
	var d FilesystemDiff
	d = MakeFilesystemDiff(left, right)
	validate(d)
	assert.NotPanics(t, func() {
		_ = d.String()
	})
}

func TestMakeFilesystemDiff_IncrementalSnapshots(t *testing.T) {

	l := fsvlist

	// basic functionality
	doTest(t, l("@a,1", "@b,2"), l("@a,1", "@b,2", "@c,3", "@d,4"), func(d FilesystemDiff) {
		assert.Equal(t, l("@b,2", "@c,3", "@d,4"), d.IncrementalPath)
	})

	// no common ancestor
	doTest(t, l(), l("@a,1"), func(d FilesystemDiff) {
		assert.Nil(t, d.IncrementalPath)
		assert.EqualValues(t, d.Conflict, ConflictNoCommonAncestor)
		assert.Equal(t, l("@a,1"), d.MRCAPathRight)
	})
	doTest(t, l("@a,1", "@b,2"), l("@c,3", "@d,4"), func(d FilesystemDiff) {
		assert.Nil(t, d.IncrementalPath)
		assert.EqualValues(t, d.Conflict, ConflictNoCommonAncestor)
		assert.Equal(t, l("@c,3", "@d,4"), d.MRCAPathRight)
	})

	// divergence is detected
	doTest(t, l("@a,1", "@b1,2"), l("@a,1", "@b2,3"), func(d FilesystemDiff) {
		assert.Nil(t, d.IncrementalPath)
		assert.EqualValues(t, d.Conflict, ConflictDiverged)
		assert.Equal(t, l("@a,1", "@b1,2"), d.MRCAPathLeft)
		assert.Equal(t, l("@a,1", "@b2,3"), d.MRCAPathRight)
	})

	// gaps before most recent common ancestor do not matter
	doTest(t, l("@a,1", "@b,2", "@c,3"), l("@a,1", "@c,3", "@d,4"), func(d FilesystemDiff) {
		assert.Equal(t, l("@c,3", "@d,4"), d.IncrementalPath)
	})

	// no snapshots on remote
	doTest(t, l(), l(), func(d FilesystemDiff) {
		assert.Equal(t, ConflictNoCommonAncestor, d.Conflict)
		assert.EqualValues(t, []FilesystemVersion(nil), d.IncrementalPath)
		assert.EqualValues(t, l(), d.MRCAPathRight)
		assert.EqualValues(t, l(), d.MRCAPathLeft)
	})
	// equivalent to no snapshots on remote
	doTest(t, nil, l(), func(d FilesystemDiff) {
		assert.Equal(t, ConflictNoCommonAncestor, d.Conflict)
		assert.EqualValues(t, []FilesystemVersion(nil), d.IncrementalPath)
		assert.EqualValues(t, l(), d.MRCAPathRight)
		assert.EqualValues(t, l(), d.MRCAPathLeft)
	})

}

func TestMakeFilesystemDiff_BookmarksSupport(t *testing.T) {
	l := fsvlist

	// bookmarks are used
	doTest(t, l("@a,1"), l("#a,1", "@b,2"), func(d FilesystemDiff) {
		assert.Equal(t, l("#a,1", "@b,2"), d.IncrementalPath)
	})

	// boomarks are stripped from IncrementalPath (cannot send incrementally)
	doTest(t, l("@a,1"), l("#a,1", "#b,2", "@c,3"), func(d FilesystemDiff) {
		assert.Equal(t, l("#a,1", "@c,3"), d.IncrementalPath)
	})

	// test that snapshots are preferred over bookmarks in IncrementalPath
	doTest(t, l("@a,1"), l("#a,1", "@a,1", "@b,2"), func(d FilesystemDiff) {
		assert.Equal(t, l("@a,1", "@b,2"), d.IncrementalPath)
	})
	doTest(t, l("@a,1"), l("@a,1", "#a,1", "@b,2"), func(d FilesystemDiff) {
		assert.Equal(t, l("@a,1", "@b,2"), d.IncrementalPath)
	})

}
