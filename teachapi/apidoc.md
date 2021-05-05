查询指定项目属性接口
-----------

### 1. 查询指定项目属性

#### 接口功能

> 获取制定项目的分类信息

#### URL

> [http://www.api.com/index.php](www.api.com/index.php)

#### 支持格式

> JSON

#### HTTP请求方式

> GET

#### 请求参数

|参数|必选|类型|说明|
|:----- |:-------|:-----|----- |
|name |ture |string|请求的项目名 |
|type |true |int |请求项目的类型。1：类型一；2：类型二 。|

#### 返回字段

|返回字段|字段类型|说明 |
|:----- |:------|:----------------------------- |
|status | int |返回结果状态。0：正常；1：错误。 |
|company | string | 所属公司名 |
|category | string |所属类型 |

#### 接口示例

> 地址：[http://www.api.com/index.php?name=可口可乐&type=1](http://www.api.com/index.php?name=可口可乐&type=1)

{
"state": 0,
"company": "可口可乐",
"category": "饮料"
}
