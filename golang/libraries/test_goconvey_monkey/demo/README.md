# Golang单元测试实操

## 考察范围

- 单元测试基础概念
- 单元测试用例设计
- 代码覆盖率

## 考试流程
1. Fork项目代码 项目地址：https://git.code.oa.com/iCode_UnitTest/Golang_Unitest
2. 在自己Fork的项目中拉分支进行开发（分支命名使用feature/{企业微信id}/unit_test_exam `git checkout -b feature/{企业微信id}/unit_test_exam`
3. 在分支编写单元测试, 对phptool.go中的函数 Strtotime 、IsStringNumeric 、IsNumeric 做单元测试
4. 提交分支代码，并向iCode_UnitTest/Golang_Unitest/master发起Merge Request，MR标题为：` [企业微信id]提交Golang单元测试作业 `， 此时会自动触发流水线检测。
5. 发起MR通过检测即可，不需要最终合入Master。

## 检测指标

- 单元测试用例总数  `>=2`
- 单元测试成功率    `>=100%`
- 单元测试运行时长 `<= 5(s)` （不包括编译时间和模拟器启动时间等，只统计单元测试方法实际执行时间）
- 代码覆盖率 `>=80%`

## 注意事项

- 不要修改phptool.go的任何代码，否则会按失败处理。
- 自己新建文件写单元测试。
- 开启go mod : ```export GO111MODULE=on```

