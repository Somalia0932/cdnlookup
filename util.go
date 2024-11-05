package main

var countryMap map[string]string
var CityMap map[string]string

func Initlist() {
	CityMap = make(map[string]string)
	CityMap["北京市 移动"] = "211.136.18.0"
	CityMap["北京市 联通"] = "202.106.0.0"
	CityMap["北京市 电信"] = "219.141.136.0"
	
	CityMap["上海市 移动"] = "211.136.150.0"
	CityMap["上海市 联通"] = "210.22.70.0"
	CityMap["上海市 电信"] = "202.96.209.0"
	
	CityMap["天津市 移动"] = "211.137.160.0"
	CityMap["天津市 联通"] = "202.99.96.0"
	CityMap["天津市 电信"] = "219.150.32.0"
	
	CityMap["重庆市 移动"] = "218.201.21.0"
	CityMap["重庆市 联通"] = "221.5.203.0"
	CityMap["重庆市 电信"] = "61.128.192.0"
	
	CityMap["河北省-石家庄市 移动"] = "211.138.13.0"
	CityMap["河北省-石家庄市 联通"] = "202.99.160.0"
	CityMap["河北省-石家庄市 电信"] = "222.222.222.0"
	
	CityMap["山西省-太原市 移动"] = "211.138.106.0"
	CityMap["山西省-太原市 联通"] = "202.99.192.0"
	CityMap["山西省-太原市 电信"] = "59.49.49.0"
	
	CityMap["辽宁省-沈阳市 移动"] = "211.137.32.0"
	CityMap["辽宁省-沈阳市 联通"] = "202.96.64.0"
	CityMap["辽宁省-沈阳市 电信"] = "219.148.204.0"
	
	CityMap["吉林省-长春市 移动"] = "211.141.16.0"
	CityMap["吉林省-长春市 联通"] = "202.98.0.0"
	CityMap["吉林省-长春市 电信"] = "219.149.194.0"
	
	CityMap["黑龙江省-哈尔滨市 移动"] = "211.137.241.0"
	CityMap["黑龙江省-哈尔滨市 联通"] = "202.97.224.0"
	CityMap["黑龙江省-哈尔滨市 电信"] = "112.100.100.0"
	
	CityMap["江苏省-南京市 移动"] = "221.130.13.0"
	CityMap["江苏省-南京市 联通"] = "221.6.4.0"
	CityMap["江苏省-南京市 电信"] = "218.2.2.0"
	
	CityMap["浙江省-杭州市 移动"] = "211.140.13.0"
	CityMap["浙江省-杭州市 联通"] = "221.12.1.0"
	CityMap["浙江省-杭州市 电信"] = "202.101.172.0"
	
	CityMap["安徽省-合肥市 移动"] = "211.138.180.0"
	CityMap["安徽省-合肥市 联通"] = "218.104.78.0"
	CityMap["安徽省-合肥市 电信"] = "61.132.163.0"
	
	CityMap["福建省-福州市 移动"] = "211.143.181.0"
	CityMap["福建省-福州市 联通"] = "58.22.96.0"
	CityMap["福建省-福州市 电信"] = "218.85.152.0"
	
	CityMap["江西省-南昌市 移动"] = "211.141.90.0"
	CityMap["江西省-南昌市 联通"] = "220.248.192.0"
	CityMap["江西省-南昌市 电信"] = "202.101.224.0"
	
	CityMap["山东省-济南市 移动"] = "211.137.191.0"
	CityMap["山东省-济南市 联通"] = "202.102.152.0"
	CityMap["山东省-济南市 电信"] = "219.146.0.0"
	
	CityMap["河南省-郑州市 移动"] = "211.138.24.0"
	CityMap["河南省-郑州市 联通"] = "202.102.224.0"
	CityMap["河南省-郑州市 电信"] = "222.85.85.0"
	
	CityMap["湖北省-武汉市 移动"] = "211.137.58.0"
	CityMap["湖北省-武汉市 联通"] = "218.104.111.0"
	CityMap["湖北省-武汉市 电信"] = "202.103.24.0"
	
	CityMap["湖南省-长沙市 移动"] = "211.142.210.0"
	CityMap["湖南省-长沙市 联通"] = "58.20.127.0"
	CityMap["湖南省-长沙市 电信"] = "202.103.96.0"
	
	CityMap["广东省-广州市 移动"] = "211.136.20.0"
	CityMap["广东省-广州市 联通"] = "210.21.4.0"
	CityMap["广东省-广州市 电信"] = "61.144.56.0"
	
	CityMap["海南省-海口市 移动"] = "221.176.88.0"
	CityMap["海南省-海口市 联通"] = "221.11.132.0"
	CityMap["海南省-海口市 电信"] = "202.100.192.0"
	
	CityMap["四川省-成都市 移动"] = "211.137.82.0"
	CityMap["四川省-成都市 联通"] = "119.6.6.0"
	CityMap["四川省-成都市 电信"] = "61.139.2.0"
	
	CityMap["贵州省-贵阳市 移动"] = "211.139.5.0"
	CityMap["贵州省-贵阳市 联通"] = "221.13.30.0"
	CityMap["贵州省-贵阳市 电信"] = "202.98.192.0"
	
	CityMap["云南省-昆明市 移动"] = "211.139.29.0"
	CityMap["云南省-昆明市 联通"] = "221.3.131.0"
	CityMap["云南省-昆明市 电信"] = "222.172.200.0"
	
	CityMap["陕西省-西安市 移动"] = "211.137.130.0"
	CityMap["陕西省-西安市 联通"] = "221.11.1.0"
	CityMap["陕西省-西安市 电信"] = "218.30.19.0"
	
	CityMap["甘肃省-兰州市 移动"] = "218.203.160.0"
	CityMap["甘肃省-兰州市 联通"] = "221.7.34.0"
	CityMap["甘肃省-兰州市 电信"] = "202.100.64.0"
	
	CityMap["青海省-西宁市 移动"] = "211.138.75.0"
	CityMap["青海省-西宁市 联通"] = "221.207.58.0"
	CityMap["青海省-西宁市 电信"] = "202.100.128.0"
}
