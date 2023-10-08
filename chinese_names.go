package namegen

var (
	chineseMaleFirstNames = []string{
		"Bin",
		"Bo",
		"Chang",
		"Chao",
		"Chen",
		"Cheng",
		"Chuan",
		"Chun",
		"Dong",
		"Feng",
		"Gang",
		"Guo",
		"Han",
		"Hao",
		"Heng",
		"Hong",
		"Huan",
		"Hui",
		"Ji",
		"Jia",
		"Jian",
		"Jiang",
		"Jie",
		"Jin",
		"Jing",
		"Jun",
		"Junjie",
		"Kai",
		"Kang",
		"Kun",
		"Lei",
		"Li",
		"Liang",
		"Lin",
		"Long",
		"Meng",
		"Ming",
		"Mingyu",
		"Peng",
		"Qiang",
		"Qing",
		"Ran",
		"Rong",
		"Rui",
		"Sheng",
		"Shi",
		"Shuai",
		"Tao",
		"Tian",
		"Tianyu",
		"Wei",
		"Wen",
		"Xia",
		"Xian",
		"Xiang",
		"Xiao",
		"Xiaofeng",
		"Xin",
		"Xing",
		"Xiong",
		"Xu",
		"Xue",
		"Yan",
		"Yang",
		"Yi",
		"Yifan",
		"Yifeng",
		"Ying",
		"Yong",
		"Yu",
		"Yuan",
		"Yue",
		"Yuhang",
		"Yun",
		"Zhang",
		"Zhao",
		"Zhe",
		"Zhen",
		"Zheng",
		"Zhi",
		"Zhihao",
		"Zhong",
		"Zihao",
		"Ziyang",
		"Ziyi",
		"Ziyu",
		"Zong",
	}

	chineseFemaleFirstNames = []string{
		"Fang",
		"Fangfang",
		"Fangyan",
		"Fenfen",
		"Feng",
		"Hailing",
		"Hua",
		"Huan",
		"Hui",
		"Huihui",
		"Huijun",
		"Huilin",
		"Jia",
		"Jian",
		"Jianmei",
		"Jiayi",
		"Jing",
		"Jinghua",
		"Jingjing",
		"Jingwen",
		"Jingyi",
		"Lan",
		"Lanlan",
		"Li",
		"Lihua",
		"Lili",
		"Lina",
		"Ling",
		"Lingling",
		"Lingyan",
		"Lingyun",
		"Lirong",
		"Lixin",
		"Liying",
		"Mei",
		"Meifang",
		"Meimei",
		"Meng",
		"Min",
		"Ming",
		"Mingming",
		"Minmin",
		"Ning",
		"Qian",
		"Qianhui",
		"Qianmei",
		"Qianqian",
		"Qianyu",
		"Qing",
		"Qingling",
		"Qingmei",
		"Qingqing",
		"Qingyuan",
		"Rong",
		"Rui",
		"Ruilin",
		"Ruyi",
		"Shanshan",
		"Wei",
		"Xia",
		"Xiaofang",
		"Xiaofei",
		"Xiaofeng",
		"Xiaohong",
		"Xiaohui",
		"Xiaojie",
		"Xiaojing",
		"Xiaojuan",
		"Xiaolan",
		"Xiaolei",
		"Xiaoli",
		"Xiaoliang",
		"Xiaoling",
		"Xiaolong",
		"Xiaomei",
		"Xiaomin",
		"Xiaona",
		"Xiaoqin",
		"Xiaoqiong",
		"Xiaorong",
		"Xiaorui",
		"Xiaotao",
		"Xiaowei",
		"Xiaowen",
		"Xiaoxi",
		"Xiaoxia",
		"Xiaoxiao",
		"Xiaoxin",
		"Xiaoxiu",
		"Xiaoxuan",
		"Xiaoxue",
		"Xiaoyan",
		"Xiaoyi",
		"Xiaoying",
		"Xiaoyue",
		"Xiaoyun",
		"Xiaozhen",
		"Xiaozhong",
		"Xin",
		"Xiu",
		"Xiufen",
		"Xiuying",
		"Xue",
		"Yan",
		"Yanfang",
		"Yanfen",
		"Yanhua",
		"Yanjuan",
		"Yanli",
		"Yanmei",
		"Yanping",
		"Yanqing",
		"Yanrong",
		"Yanxi",
		"Yanxia",
		"Yanxiang",
		"Yanyan",
		"Yanying",
		"Yanyu",
		"Yaxin",
		"Yaxuan",
		"Yi",
		"Ying",
		"Yingying",
		"Yiwei",
		"Yong",
		"Yu",
		"Yuanyuan",
		"Yue",
		"Yuemei",
		"Yufeng",
		"Yuhong",
		"Yujiao",
		"Yujie",
		"Yumei",
		"Yun",
		"Yuqin",
		"Yuxin",
		"Zhi",
		"Zhu",
	}
	chineseLastNames = []string{
		"Cao",
		"Chang",
		"Chen",
		"Cheng",
		"Chu",
		"Cui",
		"Dai",
		"Deng",
		"Du",
		"Fang",
		"Feng",
		"Fu",
		"Gao",
		"Gong",
		"Guo",
		"Han",
		"Hao",
		"He",
		"Hou",
		"Hu",
		"Hua",
		"Huang",
		"Jiang",
		"Kong",
		"Lei",
		"Li",
		"Liang",
		"Liao",
		"Lin",
		"Liu",
		"Lu",
		"Ma",
		"Mo",
		"Pan",
		"Peng",
		"Qian",
		"Qin",
		"Shen",
		"Shi",
		"Song",
		"Su",
		"Sun",
		"Tang",
		"Tao",
		"Tian",
		"Wang",
		"Wei",
		"Wu",
		"Xiang",
		"Xiao",
		"Xie",
		"Xiong",
		"Xu",
		"Yang",
		"Yao",
		"Yin",
		"You",
		"Yu",
		"Yuan",
		"Zeng",
		"Zhang",
		"Zhao",
		"Zhou",
		"Zhu",
		"Zou",
	}
)
