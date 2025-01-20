package utils

import "sort"

type TreeItemFeature interface {
	GetId() uint64
	GetParentId() uint64
	SetChildren(children []TreeItemFeature)

	GetChildren() []TreeItemFeature
}

// TransformListToTreeData 将分类的平铺结构转成树形结构
func TransformListToTreeData[T TreeItemFeature](list []T) []T {
	size := len(list)
	if list == nil || size == 0 {
		return []T{}
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].GetParentId() < list[j].GetParentId()
	})
	treeData := make([]T, 0)
	dataMap := make(map[uint64]T, size)

	for _, value := range list {
		dataMap[value.GetId()] = value
	}
	for _, value := range dataMap {
		if value.GetParentId() > 0 {
			parentNode := dataMap[value.GetParentId()]
			parentNode.SetChildren(append(parentNode.GetChildren(), value))
		} else {
			treeData = append(treeData, value)
		}
	}
	return treeData
}

func SoreTreeData[T TreeItemFeature](treeData []T, sortFunc func(i, j TreeItemFeature) bool) []T {
	if len(treeData) == 0 {
		return treeData
	}
	for index := range treeData {
		treeData[index].SetChildren(SoreTreeData(treeData[index].GetChildren(), sortFunc))
	}
	sort.Slice(treeData, func(i, j int) bool {
		return sortFunc(treeData[i], treeData[j])
	})
	return treeData
}
