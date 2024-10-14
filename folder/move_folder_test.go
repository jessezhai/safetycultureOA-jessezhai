package folder_test

import (
	"testing"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/georgechieng-sc/interns-2022/folder"
)

func Test_folder_MoveFolder(t *testing.T) {

	folders := []folder.Folder{
		{Name: "alpha", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "alpha"},
		{Name: "bravo", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "alpha.bravo"},
		{Name: "charlie", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "alpha.charlie"},
		{Name: "delta", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "delta"},
		{Name: "foxtrot", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "delta.foxtrot"},
		{Name: "gamma", OrgId: uuid.FromStringOrNil(folder.DefaultOrgID), Paths: "alpha.gamma"},
		{Name: "zulu", OrgId: uuid.Must(uuid.NewV4()), Paths: "zulu"},
	}

	folderDriver := folder.NewDriver(folders)

	tests := []struct {
		name          string
		source        string
		destination   string
		expectedError string
		expectedPath  string
	}{
		{"Valid Move", "bravo", "charlie", "", "alpha.charlie.bravo"},
		{"Move to Itself", "bravo", "bravo", "Cannot move a folder to itself", ""},
		{"Move to Child of Itself", "bravo", "charlie", "Cannot move a folder to a child of itself", ""},
		{"Move to Different Organization", "bravo", "zulu", "Cannot move a folder to a different organization", ""},
		{"Invalid Source Folder", "invalid_folder", "delta", "Source folder does not exist", ""},
		{"Invalid Destination Folder", "bravo", "invalid_folder", "Destination folder does not exist", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updatedFolders, err := folderDriver.MoveFolder(tt.source, tt.destination)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err.Error())
				return
			}

			assert.NoError(t, err)
			for _, f := range updatedFolders {
				if f.Name == tt.source {
					assert.Equal(t, tt.expectedPath, f.Paths, "Path of the moved folder is incorrect")
					return
				}
			}
			assert.Fail(t, "folder not found")
		})
	}
}
