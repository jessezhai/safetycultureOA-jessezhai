package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()
	// Mock data
	orgID1, _ := uuid.NewV4()
	orgID2, _ := uuid.NewV4()
	folder1 := folder.Folder{OrgId: orgID1, Name: "Folder1", Paths: "Folder1"}
	folder2 := folder.Folder{OrgId: orgID1, Name: "Folder2", Paths: "Folder1.Folder2"}
	folder3 := folder.Folder{OrgId: orgID2, Name: "Folder3", Paths: "Folder3"}

	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		// TODO: your tests here
		{
			name:    "Get folders for OrgID 1, positive test case.",
			orgID:   orgID1,
			folders: []folder.Folder{folder1, folder2, folder3},
			want:    []folder.Folder{folder1, folder2},
		},
		{
			name:    "Get folders for OrgID 2, positive test case.",
			orgID:   orgID2,
			folders: []folder.Folder{folder1, folder2, folder3},
			want:    []folder.Folder{folder3},
		},
		{
			name:    "No folders for OrgID 3, negative test case.",
			orgID:   uuid.Must(uuid.NewV4()),
			folders: []folder.Folder{folder1, folder2, folder3},
			want:    []folder.Folder{},
		},
		{
			name:    "Empty folder list, negative test case.",
			orgID:   orgID1,
			folders: []folder.Folder{},
			want:    []folder.Folder{},
		},
		{
			name:    "Invalid OrgID (nil), edge test case.",
			orgID:   uuid.Nil,
			folders: []folder.Folder{folder1, folder2, folder3},
			want:    []folder.Folder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get := f.GetFoldersByOrgID(tt.orgID)
			assert.Equal(t, tt.want, get)
		})
	}
}

func Test_folder_GetAllChildFolders(t *testing.T) {
	t.Parallel()

	// Mock data
	orgID1, _ := uuid.NewV4()

	folder1 := folder.Folder{OrgId: orgID1, Name: "Folder1", Paths: "Folder1"}
	folder2 := folder.Folder{OrgId: orgID1, Name: "Folder2", Paths: "Folder1.Folder2"}
	folder3 := folder.Folder{OrgId: orgID1, Name: "Folder3", Paths: "Folder1.Folder2.Folder3"}

	tests := []struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		parent  string
		want    []folder.Folder
	}{
		{
			name:    "Get all child folders for Folder1, positive test case.",
			orgID:   orgID1,
			folders: []folder.Folder{folder1, folder2, folder3},
			parent:  "Folder1",
			want:    []folder.Folder{folder2, folder3},
		},
		{
			name:    "Get all child folders for Folder2, positive test case.",
			orgID:   orgID1,
			folders: []folder.Folder{folder1, folder2, folder3},
			parent:  "Folder2",
			want:    []folder.Folder{folder3},
		},
		{
			name:    "Folder does not exist, negative test case.",
			orgID:   orgID1,
			folders: []folder.Folder{folder1, folder2, folder3},
			parent:  "NonExistentFolder",
			want:    nil,
		},
		{
			name:    "Empty folder list, negative test case.",
			orgID:   orgID1,
			folders: []folder.Folder{},
			parent:  "Folder1",
			want:    nil,
		},
		{
			name:    "Parent folder path only partially matches child folder, negative test case",
			orgID:   orgID1,
			folders: []folder.Folder{
				{OrgId: orgID1, Name: "FolderA", Paths: "FolderA"},
				{OrgId: orgID1, Name: "FolderB", Paths: "FolderAB"},
			},
			parent: "FolderA",
			want:   []folder.Folder{},
		},
		{
			name:    "Parent folder exists but has no children, negative test case",
			orgID:   orgID1,
			folders: []folder.Folder{folder1},
			parent:  "Folder1",
			want:    []folder.Folder{},
		},
		{
			name:    "Invalid OrgID (nil) with valid parent, edge test case",
			orgID:   uuid.Nil,
			folders: []folder.Folder{folder1, folder2, folder3},
			parent:  "Folder1",
			want:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			got := f.GetAllChildFolders(tt.orgID, tt.parent)
			assert.Equal(t, tt.want, got)
		})
	}
}	
