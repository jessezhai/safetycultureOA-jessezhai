package folder

import (
	"errors"
	"strings"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {

	folders := f.folders

	// Find source folder and subtree
	var sourceFolder Folder
	var sourceSubtree []Folder
	sourceIndex := -1
	for i, folder := range folders {
		if folder.Name == name {
			sourceFolder = folder
			sourceIndex = i
			break
		}
	}
	if sourceIndex == -1 {
		return nil, errors.New("Source folder does not exist")
	}

	// Find all child folders of source folder
	for _, folder := range folders {
		if strings.HasPrefix(folder.Paths, sourceFolder.Paths+".") {
			sourceSubtree = append(sourceSubtree, folder)
		}
	}

	// Find the destination folder
	var destFolder Folder
	destIndex := -1
	for i, folder := range folders {
		if folder.Name == dst {
			destFolder = folder
			destIndex = i
			break
		}
	}
	if destIndex == -1 {
		return nil, errors.New("Destination folder does not exist")
	}

	if sourceFolder.OrgId != destFolder.OrgId {
		return nil, errors.New("Cannot move a folder to a different organization")
	}

	if sourceFolder.Name == destFolder.Name {
		return nil, errors.New("Cannot move a folder to itself")
	}

	if strings.HasPrefix(sourceFolder.Paths, destFolder.Paths) {
		return nil, errors.New("Cannot move a folder to a child of itself")
	}

	// Update paths of source folder and children
	newPath := destFolder.Paths + "." + sourceFolder.Name
	sourceFolder.Paths = newPath
	for i := range sourceSubtree {
		sourceSubtree[i].Paths = strings.Replace(sourceSubtree[i].Paths, sourceFolder.Paths, newPath, 1)
	}

	// Remove source folder and subtree from the original location
	folders = append(folders[:sourceIndex], folders[sourceIndex+1:]...)
	for _, childFolder := range sourceSubtree {
		for i, folder := range folders {
			if folder.Paths == childFolder.Paths {
				folders = append(folders[:i], folders[i+1:]...)
				break
			}
		}
	}

	// Add moved folder and its subtree to the new destination
	folders = append(folders, sourceFolder)
	folders = append(folders, sourceSubtree...)

	return folders, nil
}
