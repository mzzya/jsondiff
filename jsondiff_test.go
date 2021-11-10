package jsondiff

import (
	"reflect"
	"testing"
)

func TestDiff(t *testing.T) {
	type args struct {
		json1 string
		json2 string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{name: "simple test", args: args{json1: `{"int":1,"string":"string","float":1.1,"array":[1,2,3],"object":{"int":1,"string":"string"},"objectAry":[{"int":1,"string":"string"}]}`, json2: `{"int":12,"string":"string2","float":1.12,"array":[12,22,32],"object":{"int":1,"string":"string2"}}`}, want: nil, wantErr: true},
		{name: "simple test2", args: args{json1: `{"data":{"projectId":203,"projectCode":"CBHB","projectName":"渤采银行","thirdPartyId":"0","memberCode":"","secretKey":"Aa123456","orderCodeExtendState":0,"orderCodeExtendField":"","isEnableCostMatch":true,"isEnableCostSpecial":false,"isEnableDefaultCost":true,"isEnableAreaMatch":true,"autoMatchExternalCustomerCode":false,"areaMatchField":"DistrictCode","unitMapping":"什么鬼","salesContact":null,"partnerContact":"","noticeAddress":"11@example.com","isEnableHttps":false,"isSendProductName":false,"version":"1.0","remark":"cip","isRelease":false,"operationType":1,"priority":17,"projectType":0},"isSuccess":true,"code":"0","message":""}`, json2: `{"data":{"id":194,"projectId":203,"thirdPartyId":"0","memberCode":"","secretKey":"Aa123456","orderCodeExtendState":0,"orderCodeExtendField":"","areaMatchField":"DistrictCode","unitMapping":"什么鬼","salesContact":null,"partnerContact":"","noticeAddress":"tansi@colipu.com","version":"1.0","remark":"cip","enable":null,"release":null,"creationTime":"2020-04-15 13:25:00","creatorId":1,"creator":"cip","modificationTime":"2020-05-09 17:45:57","modifierId":1,"modifier":"cip","platformAddress":null,"interfacePovider":null,"salesTeam":null,"validityDate":"2021-12-30 00:00:00","signContractMethod":3,"sales":null,"operationType":1,"priority":17,"projectType":0,"autoMatchExternalCustomerCode":false,"isEnableCostMatch":false,"isEnableCostSpecial":null,"isEnableDefaultCost":null,"isEnableAreaMatch":null,"isEnableHttps":null,"isSendProductName":null},"isSuccess":true,"message":""}`}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Diff(tt.args.json1, tt.args.json2, false)
			if (err != nil) != tt.wantErr {
				t.Errorf("Diff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				// t.Errorf("Diff() = %v, want %v", got, tt.want)
			}
		})
	}
}

var json1 = `{"StatusCode":"SUCCESS","Message":null,"RequestId":"0HM3VJMG2NREK:00000029","Data":[{"ReceiverId":1806979,"ShipTypeId":9,"CompanyFullName":"","CompanyBriefName":"","ContactName":"CHJ","ContactPhone":"021-31216888","ContactMobile":"15221987518","Fax":"021-1111111","EMail":"qinsicheng@colipu.com","Zip":"200000","Address":"测试请勿配送,谢谢","Longitude":121.556689,"Latitude":30.9198,"Remark":"","CreateUserId":0,"CreateTime":"2016-08-07T12:56:52.85+08:00","UpdateUserId":0,"UpdateTime":"2016-08-07T12:56:52.85+08:00","Status":"A","CustomerId":187288,"ProvinceId":2,"ProvinceName":"上海","CityId":3,"CityName":"上海市","DistrictId":20,"DistrictName":"奉贤区","IsDefault":"N","AccountID":0,"DefaultWarehouseId":0,"DefaultLogicalWarehouseId":0},{"ReceiverId":2000061,"ShipTypeId":0,"CompanyFullName":"","CompanyBriefName":"","ContactName":"323233","ContactPhone":"213124","ContactMobile":"","Fax":"","EMail":"","Zip":"232142","Address":"412414324","Longitude":121.481238,"Latitude":31.213348,"Remark":"","CreateUserId":100831,"CreateTime":"2017-12-24T18:01:38.743+08:00","UpdateUserId":100874,"UpdateTime":"2017-12-24T18:02:20.41+08:00","Status":"A","CustomerId":187288,"ProvinceId":2,"ProvinceName":"上海","CityId":3,"CityName":"上海市","DistrictId":5,"DistrictName":"卢湾区","IsDefault":"N","AccountID":0,"DefaultWarehouseId":0,"DefaultLogicalWarehouseId":0},{"ReceiverId":2339607,"ShipTypeId":0,"CompanyFullName":"","CompanyBriefName":"","ContactName":"华南","ContactPhone":"111","ContactMobile":"","Fax":"","EMail":"","Zip":"111111","Address":"广州","Longitude":113.220176,"Latitude":23.446661,"Remark":"","CreateUserId":100874,"CreateTime":"2018-09-30T11:16:47.227+08:00","UpdateUserId":100913,"UpdateTime":"2020-02-14T10:46:34.72+08:00","Status":"A","CustomerId":187288,"ProvinceId":269,"ProvinceName":"广东省","CityId":280,"CityName":"广州市","DistrictId":281,"DistrictName":"花都区","IsDefault":"Y","AccountID":0,"DefaultWarehouseId":0,"DefaultLogicalWarehouseId":0},{"ReceiverId":2339606,"ShipTypeId":0,"CompanyFullName":"","CompanyBriefName":"","ContactName":"华北","ContactPhone":"11111","ContactMobile":"","Fax":"","EMail":"","Zip":"111111","Address":"北京","Longitude":116.521695,"Latitude":39.958953,"Remark":"","CreateUserId":100874,"CreateTime":"2018-09-30T11:15:43.3+08:00","UpdateUserId":100913,"UpdateTime":"2020-02-14T10:46:34.72+08:00","Status":"A","CustomerId":187288,"ProvinceId":23,"ProvinceName":"北京","CityId":274,"CityName":"北京市","DistrictId":275,"DistrictName":"朝阳区","IsDefault":"N","AccountID":0,"DefaultWarehouseId":0,"DefaultLogicalWarehouseId":0},{"ReceiverId":2339602,"ShipTypeId":0,"CompanyFullName":"","CompanyBriefName":"","ContactName":"李丹dandandandandan","ContactPhone":"","ContactMobile":"18792390491","Fax":"","EMail":"","Zip":"200030","Address":"古美路1528号A2幢12层","Longitude":121.404331,"Latitude":31.170375,"Remark":"","CreateUserId":100874,"CreateTime":"2018-09-30T11:14:07.617+08:00","UpdateUserId":100913,"UpdateTime":"2020-02-14T10:46:34.72+08:00","Status":"A","CustomerId":187288,"ProvinceId":2,"ProvinceName":"上海","CityId":3,"CityName":"上海市","DistrictId":7,"DistrictName":"徐汇区","IsDefault":"N","AccountID":0,"DefaultWarehouseId":0,"DefaultLogicalWarehouseId":0}],"BusinessCodeMessage":null}`
var json2 = `{"StatusCode":"SUCCESS","Message":null,"RequestId":"2ed7a3bd-49f9-45d7-a07b-5f3f85aa1622","Data":[{"ReceiverId":1806979,"ShipTypeId":9,"CompanyFullName":"","CompanyBriefName":"","ContactName":"CHJ","ContactPhone":"021-31216888","ContactMobile":"15221987518","Fax":"021-1111111","Zip":"200000","Address":"测试请勿配送,谢谢","Longitude":121.556689,"Latitude":30.9198,"Remark":"","CreateUserId":0,"CreateTime":"2016-08-07T12:56:52.850+08:00","UpdateUserId":0,"UpdateTime":"2016-08-07T12:56:52.850+08:00","Status":"A","CustomerId":187288,"ProvinceId":2,"CityId":3,"DistrictId":20,"ProvinceName":"上海","CityName":"上海市","DistrictName":"奉贤区","IsDefault":"N","DefaultWarehouseId":0,"DefaultLogicalWarehouseId":0,"Email":"qinsicheng@colipu.com","OperatorID":null,"AccountID":null},{"ReceiverId":2000061,"ShipTypeId":0,"CompanyFullName":"","CompanyBriefName":"","ContactName":"323233","ContactPhone":"213124","ContactMobile":"","Fax":"","Zip":"232142","Address":"412414324","Longitude":121.481238,"Latitude":31.213348,"Remark":"","CreateUserId":100831,"CreateTime":"2017-12-24T18:01:38.743+08:00","UpdateUserId":100874,"UpdateTime":"2017-12-24T18:02:20.410+08:00","Status":"A","CustomerId":187288,"ProvinceId":2,"CityId":3,"DistrictId":5,"ProvinceName":"上海","CityName":"上海市","DistrictName":"卢湾区","IsDefault":"N","DefaultWarehouseId":0,"DefaultLogicalWarehouseId":0,"Email":"","OperatorID":null,"AccountID":null},{"ReceiverId":2339607,"ShipTypeId":0,"CompanyFullName":"","CompanyBriefName":"","ContactName":"华南","ContactPhone":"111","ContactMobile":"","Fax":"","Zip":"111111","Address":"广州","Longitude":113.220176,"Latitude":23.446661,"Remark":"","CreateUserId":100874,"CreateTime":"2018-09-30T11:16:47.227+08:00","UpdateUserId":100913,"UpdateTime":"2020-02-14T10:46:34.720+08:00","Status":"A","CustomerId":187288,"ProvinceId":269,"CityId":280,"DistrictId":281,"ProvinceName":"广东省","CityName":"广州市","DistrictName":"花都区","IsDefault":"Y","DefaultWarehouseId":129,"DefaultLogicalWarehouseId":10031,"Email":"","OperatorID":null,"AccountID":null},{"ReceiverId":2339606,"ShipTypeId":0,"CompanyFullName":"","CompanyBriefName":"","ContactName":"华北","ContactPhone":"11111","ContactMobile":"","Fax":"","Zip":"111111","Address":"北京","Longitude":116.521695,"Latitude":39.958953,"Remark":"","CreateUserId":100874,"CreateTime":"2018-09-30T11:15:43.300+08:00","UpdateUserId":100913,"UpdateTime":"2020-02-14T10:46:34.720+08:00","Status":"A","CustomerId":187288,"ProvinceId":23,"CityId":274,"DistrictId":275,"ProvinceName":"北京","CityName":"北京市","DistrictName":"朝阳区","IsDefault":"N","DefaultWarehouseId":129,"DefaultLogicalWarehouseId":10031,"Email":"","OperatorID":null,"AccountID":null},{"ReceiverId":2339602,"ShipTypeId":0,"CompanyFullName":"","CompanyBriefName":"","ContactName":"李丹dandandandandan","ContactPhone":"","ContactMobile":"18792390491","Fax":"","Zip":"200030","Address":"古美路1528号A2幢12层","Longitude":121.404331,"Latitude":31.170375,"Remark":"","CreateUserId":100874,"CreateTime":"2018-09-30T11:14:07.617+08:00","UpdateUserId":100913,"UpdateTime":"2020-02-14T10:46:34.720+08:00","Status":"A","CustomerId":187288,"ProvinceId":2,"CityId":3,"DistrictId":7,"ProvinceName":"上海","CityName":"上海市","DistrictName":"徐汇区","IsDefault":"N","DefaultWarehouseId":129,"DefaultLogicalWarehouseId":10031,"Email":"","OperatorID":null,"AccountID":null}],"BusinessCodeMessage":null}`

var json1Bytes = []byte(json1)
var json2Bytes = []byte(json2)

func BenchmarkDiff(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Diff(json1, json2, false)
	}
}

func BenchmarkDiffBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DiffBytes(json1Bytes, json2Bytes, false)
	}
}
