package module

var (
    UserApi *UserMgr
    CraftApi *CraftMgr
    RankApi *RankMgr
)

func Init() {
    UserApi = NewUserMgr()
    CraftApi = NewCraftMgr()
    RankApi = NewRankMgr()
}
