package taodb

type Index map[string]*position

// TODO: 这里我有个想法需要验证：由于文件分已经被merge和没有被merge两个区，那么需要每次搜索两个区寻找fileId。要在postion存在哪个区吗？
// 需要测试存和不存对内存、耗时的影响，暂时不存
type position struct {
	fileId uint32
	vsz    uint32
	vpos   uint32
	timestamp uint32
}

type TempIndex map[string]*vPosition

type vPosition struct {
	fileId uint32
	vsz uint32
	vpos uint32
	timestamp uint32
	deleted bool // 判断是否已经删除
}

