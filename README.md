# DREP声誉系统算法库



## 一、DREP声誉系统简介：


针对传统积分体系与现今用户需求的矛盾——透明度差、可用性差等痛点，DREP通过区块链、去中心化数字ID、声誉协议等因素，为商家和用户提供数据高度安全、过程公开透明，方便使用的reputation生态系统。这个系统将为众多合作平台提供具有真正使用价值的声誉，以此确定可以交易token的价值。与此同时，高声誉用户在相应平台能获得较大的影响力/优惠等优势，鼓励用户提高声誉。

区块链+声誉系统对传统的积分系统是一个较大的改进与提升，体现在：
- **声誉系统能够真正反映用户行为的有效性。** 目前互联网积分系统中，对于除了登录等简单机制之外的用户行为，往往很难真正得知用户的行为质量，同时还存在刷分等不利于积分体系正常运行的行为。在声誉系统下，用户进行评论等复杂声誉活动的时候，需要经过他人的声誉评价等更加公平有效的机制方能获得声誉，能够真正反映用户的评论质量。同时举报等机制也必须在抵押自身声誉的情况下进行，既避免了恶意举报导致的人人自危的现象，也保障用户举报的动力。
- **区块链能够有效减少作弊可能。** 用户任何获取声誉的行为，均在区块链上有所记录。区块链具有不可篡改的特性和时间戳这一利器，这样如果用户发生作弊的行径，在区块链上很容易查到相应的记录。同时，可以将刷分的行为更好的发现，调整相应的声誉系统以减少低层次的获得声誉。
- **声誉系统能够赋予积分以价值。** 由于声誉获得不容篡改，这样获得的声誉具有真正的价值。当投射到积分上时，可以通过DREP侧链发行相应的token，token可以进行用户在同平台甚至跨平台跨链交易，从而使用户真正将积分使用起来。商家之间可以调整token的汇率和兑换奖励，整合各个分散的积分，使用户能够提高粘附性的同时换取真正想要的奖励。

## 二、Drep声誉系统功能与接入方法：


#### 1、给用户分配声誉ID和对用户分组

调用GetProfile方法，传入参数商家平台号（platformID），和用户在该商家平台上的用户唯一标识（UID），返回用户在DREP声誉系统中的声誉ID（repID）与所属分组的类别（groupID)。

```
var platformID string = "taobao"
var UID string = "me"
var profile map[string] interface{} = GetProfile(platformID, UID)
// profile {
//    repID: 1234,
//    groupID: 12
// }
```

#### 2、在DREP声誉系统中注册用户

调用RegisterUser方法，传入商家平台号platformID，用户的repID与用户所属组别groupID，完成对用户在DREP声誉系统中的注册。

```
var platformID string = "taobao"
var repID string = "1234"
var groupID int64 = 12
RegisterUser(platformID, repID, groupID)
```

#### 3、将商家平台上用户每天的积分获取情况导入DREP声誉系统中

调用AddGain方法，传入商家平台号platformID，以及一组用户积分数据，其中用户积分数据包括用户repID，日期Day，和用户在日期当天获得的积分数Gain。调用完成后，这组用户数据将被记录在DREP声誉系统中，用于后面的声誉结算。

```
var platformID string = "taobao"
var increments []map[string] interface = make([]map[string] interface{}, 2)
inc1 := make(map[string] interface{})
inc1["RepID"] = "1234"
inc1["Day"] = 1
inc1["Gain"] = 20
increments[0] = inc1
inc2 := make(map[string] interface{})
inc2["RepID"] = "1234"
inc2["Day"] = 2
inc2["Gain"] = 30
increments[1] = inc2
AddGain(platformID, increments)
```

#### 4、结算用户的声誉

调用LiquidateRep方法，传入商家平台号platformID，一组用户repIDs，以及截止日期until，DREP声誉系统将自动结算所有repIDs中的用户截止到until当天前每天的声誉获得情况，并根据结算声誉发放相应的Token奖励。

```
var platformID string = "taobao"
var repIDs []string = []string{"1234", "2234", "3234", "4234"}
var until int = 2
LiquidateRep(platformID, repIDs, until)
```

#### 5、按组别结算用户的声誉

调用LiquidateRepByGroup方法，传入商家平台号platformID，组别groupID，以及截止日期until，DREP声誉系统将自动结算所有属于groupID对应组中的用户截止到until当天前每天的声誉获得情况，并根据结算声誉发放相应的Token奖励。

```
var platformID string = "taobao"
var groupID int64 = 12
var until int = 2
LiquidateRepByGroup(platformID, groupID, until)
```
