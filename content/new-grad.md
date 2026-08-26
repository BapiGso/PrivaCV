<style>
/* ===========================================================
   Engineering — 复刻 r/EngineeringResumes 那套官方推荐排版
   （RenderCV 里叫 engineeringresumes 主题）

   原则就一条：把所有装饰都删掉，只留信息密度和机器可读性。
   零配色、零图标、单栏、正文不小于 10pt、行距压到极限。
   投大厂官网、校招系统、猎头渠道用这个最稳。
   =========================================================== */

body {
    padding: 12mm 14mm;
    font-family: Inter, "Source Sans 3", "Helvetica Neue", Helvetica, Arial,
                 "Segoe UI", Roboto, "PingFang SC", "Source Han Sans SC",
                 "Microsoft YaHei", sans-serif;
    font-size: 10pt;
    line-height: 1.34;
    color: #000;
}

h1 {
    font-size: 17pt;
    font-weight: 700;
    line-height: 1.15;
    letter-spacing: 0;
    margin: 0 0 1mm;
}

h1 + p {
    color: #000;
    font-size: 9.4pt;
    line-height: 1.4;
    margin: 0 0 3.8mm;
}

h2 {
    font-size: 10.4pt;
    font-weight: 700;
    line-height: 1.25;
    text-transform: uppercase;
    letter-spacing: 0.05em;
    padding-bottom: 0.4mm;
    border-bottom: 0.9pt solid #000;
    margin: 3.8mm 0 1.2mm;
}

h2:first-of-type { margin-top: 0; }

h3 {
    font-size: 10pt;
    font-weight: 700;
    line-height: 1.3;
    margin: 2.4mm 0 0;
}

h4 {
    font-size: 9.6pt;
    font-weight: 400;
    color: #000;
    line-height: 1.35;
    margin: 0.4mm 0 0.8mm;
}

/* 时间和地点：斜体写在标题末尾，浮到右边（这套模板里一律用正体） */
h3 em, h4 em {
    float: right;
    font-style: normal;
    font-weight: 400;
    font-variant-numeric: tabular-nums;
}

p { margin: 0 0 0.8mm; }

/* 密排：项目符号浅缩进、条目之间几乎不留白 */
ul { padding-left: 4.2mm; margin: 0.6mm 0 1mm; }
li { margin: 0 0 0.4mm; }
li::marker { color: #000; font-size: 0.8em; }

a { color: inherit; text-decoration: underline; text-underline-offset: 1.4pt; text-decoration-thickness: 0.4pt; }

/* ATS 会把带底色的行内代码读成普通文本，这里干脆不做样式 */
code {
    font-family: inherit;
    font-size: 1em;
}

blockquote {
    border: 0;
    padding: 0;
    margin: 0 0 0.8mm;
    color: #000;
}
</style>

# 王思齐

计算机科学与技术 · 2026 届本科应届生 ⋄ 求职意向：后端开发 ⋄ 成都
138-0000-0000 ⋄ wangsiqi@example.com ⋄ [github.com/example](https://github.com/example)

## 教育背景

### 电子科技大学 *2022.09 - 2026.06*
#### 计算机科学与技术 · 本科 *GPA 3.82/4.0，专业排名 8/210*

核心课程：数据结构与算法、操作系统、计算机网络、数据库系统、分布式系统、编译原理

## 实习经历

### 某互联网公司 · 交易中台 *2025.06 - 2025.09*
#### 后端开发实习生 *成都*

- 独立完成优惠券核销接口的重构，用批量查询替换循环单查，接口 P95 从 620ms 降到 90ms。
- 为核心链路补充 47 个单元测试，把该模块的分支覆盖率从 31% 提升到 78%，上线后该模块零 P2 以上事故。
- 参与一次线上问题排查，定位到 Redis 大 key 导致的集群倾斜，提出并落地分片方案。

### 某创业公司 *2024.07 - 2024.09*
#### 研发实习生 *远程*

- 负责内部工单系统的后端接口开发，独立交付 12 个接口并完成联调上线。
- 编写数据迁移脚本，把 8 万条历史工单从 MongoDB 迁移到 PostgreSQL，零数据丢失。

## 项目经历

### 分布式短链服务 *2025.03 - 2025.05*
#### 个人项目 · Go + Redis + MySQL *[github.com/example/shorturl](https://github.com/example/shorturl)*

- 用发号器 + Base62 编码生成短链，单机压测 QPS 1.2 万，P99 < 15ms。
- 实现布隆过滤器拦截无效查询，缓存穿透场景下数据库压力下降 94%。
- 完整的 Docker Compose 部署与 Prometheus 监控面板，README 含压测报告。

### 校园二手书交易平台 *2024.10 - 2024.12*
#### 团队 4 人 · 担任后端负责人 *Spring Boot + Vue*

- 负责用户、商品、订单三个模块的设计与实现，定义团队的接口规范与 Git 协作流程。
- 上线后校内注册用户 2,300+，累计成交 1,100 单，稳定运行一个学期。

## 技能

- 语言：Go（熟练）、Java（熟练）、Python、C++、SQL
- 后端：Gin、Spring Boot、MySQL、Redis、Kafka、gRPC
- 工具：Git、Docker、Linux、Postman、Prometheus

## 获奖与其他

- 2024 年 ACM-ICPC 亚洲区域赛（成都站）铜奖
- 2024 年全国大学生计算机设计大赛 省级一等奖
- 校级一等奖学金（2023、2024、2025 连续三年）
- 英语 CET-6 561 分，可无障碍阅读英文技术文档
