package folder

import (
	"fmt"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {

    folders := f.folders

    res := []Folder{}

    // Get parent folder path
    var parentPath string
    foundOrg := false
    for _, folder := range folders {
        if folder.OrgId == orgID {
            foundOrg = true 
            if folder.Name == name {
                parentPath = folder.Paths
                break
            }
        }
    }

    // Folder does not exist
    if parentPath == "" {
        if !foundOrg {
			fmt.Println("Folder does not exist in the specified organization")
            return nil
        }
		fmt.Println("Folder does not exist")
        return nil
    }

	// Find child folders
    for _, folder := range folders {
        if folder.OrgId == orgID && len(folder.Paths) > len(parentPath) && folder.Paths[:len(parentPath)+1] == parentPath+"." {
            res = append(res, folder)
        }
    }

    return res
}
