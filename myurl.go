package myurl

//定义URL请求，方便传参
type UrlInfo struct {
	vin    string //vin码
	action string //动作，做什么，包含但不限于如下:
	//cata1,cata2,cata3,cata4,part,
	//all_epc_part_name,all_std_part_name,all_part_number
	//search_part_number,search_std_part_name,search_epc_part_name,search_illustration,illustration,imghotspot

	is_vin_filter_open            string //是否开启VIN码筛选 默认是1 开启
	cata1_code                    string //一级目录ID
	cata2_code                    string //二级目录ID
	cata3_code                    string //三级目录ID
	cata4_code                    string //四级目录ID
	query_match_type              string //搜索时的匹配模式 exact ,inexact
	query_part_number             string //搜索什么的配件号码
	query_part_name               string //搜索使用的配件名称
	query_part_name_is_safebase64 string //是否开启安全的配件名称搜索，默认不开，直接用名称搜索
	query_part_name_language      string //搜索配件名时使用的语言，只支持中文和英文
	illustration_img_address      string //分解图图片地址
	cata_code                     string //目录CODE
	last_cata_code                string //末级目录CODE
	last_cata_code_level          string //在请求配件列表的时候所用，大部分情况下默认可以不写，对于某些品牌需要带该参数
	guest_ip                      string //访问接口的用户所在IP,用于内网安全
	epc_id                        string //epc_id
	js_id                         string //标准车型ID
	gonggao_no                    string //公告型号
}
