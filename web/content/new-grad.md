<style>
/* ===========================================================
   Engineering — 复刻 r/EngineeringResumes 那套官方推荐排版
   （RenderCV 里叫 engineeringresumes 主题）

   原则就一条：把所有装饰都删掉，只留信息密度和机器可读性。
   零配色、零图标、单栏、正文不小于 10pt、行距压到极限。
   投大厂官网、校招系统、猎头渠道用这个最稳。
   =========================================================== */

:root {
    --font-body: var(--font-sans);

    --fs: 10pt;
    --lh: 1.34;
    --fs-name: 17pt;
    --fs-sec: 10.4pt;

    --page-x: 14mm;
    --page-y: 12mm;
    --sp-sec: 3.8mm;
    --sp-item: 2.4mm;
    --sp-line: 0.8mm;

    --ink: #000;
    --muted: #000;
    --rule: #000;
    --tint: transparent;
}

h1 {
    font-weight: 700;
    letter-spacing: 0;
    margin: 0 0 1mm;
}

h1 + p {
    color: var(--ink);
    font-size: calc(var(--fs) * 0.94);
    line-height: 1.4;
    margin: 0 0 var(--sp-sec);
}

h2 {
    font-weight: 700;
    text-transform: uppercase;
    letter-spacing: 0.05em;
    padding-bottom: 0.4mm;
    border-bottom: 0.9pt solid var(--rule);
    margin: var(--sp-sec) 0 1.2mm;
}

h3 { font-weight: 700; font-size: var(--fs); }

h4 {
    font-weight: 400;
    color: var(--ink);
    font-size: calc(var(--fs) * 0.96);
}

[data-split] > .x { color: var(--ink); }

/* 密排：项目符号浅缩进、条目之间几乎不留白 */
ul { padding-left: 4.2mm; margin: 0.6mm 0 1mm; }
li { margin-bottom: 0.4mm; }
li::marker { font-size: 0.8em; }

/* ATS 会把带底色的行内代码读成普通文本，这里干脆不做样式 */
code {
    font-family: var(--font-body);
    background: none;
    padding: 0;
    font-size: 1em;
}

a { text-decoration: underline; text-underline-offset: 1.4pt; text-decoration-thickness: 0.4pt; }

blockquote {
    border: 0;
    padding: 0;
    margin: 0 0 var(--sp-line);
    color: var(--ink);
}

td:first-child { width: 24mm; font-weight: 700; }
</style>

# 王思齐

计算机科学与技术 · 2026 届本科应届生 ⋄ 求职意向：后端开发 ⋄ 成都
138-0000-0000 ⋄ wangsiqi@example.com ⋄ [github.com/example](https://github.com/example)

## 教育背景

### 电子科技大学 <span>2022.09 - 2026.06</span>
#### 计算机科学与技术 · 本科 <span>GPA 3.82/4.0，专业排名 8/210</span>

核心课程：数据结构与算法、操作系统、计算机网络、数据库系统、分布式系统、编译原理

## 实习经历

### 某互联网公司 · 交易中台 <span>2025.06 - 2025.09</span>
#### 后端开发实习生 <span>成都</span>

- 独立完成优惠券核销接口的重构，用批量查询替换循环单查，接口 P95 从 620ms 降到 90ms。
- 为核心链路补充 47 个单元测试，把该模块的分支覆盖率从 31% 提升到 78%，上线后该模块零 P2 以上事故。
- 参与一次线上问题排查，定位到 Redis 大 key 导致的集群倾斜，提出并落地分片方案。

### 某创业公司 <span>2024.07 - 2024.09</span>
#### 研发实习生 <span>远程</span>

- 负责内部工单系统的后端接口开发，独立交付 12 个接口并完成联调上线。
- 编写数据迁移脚本，把 8 万条历史工单从 MongoDB 迁移到 PostgreSQL，零数据丢失。

## 项目经历

### 分布式短链服务 <span>2025.03 - 2025.05</span>
#### 个人项目 · Go + Redis + MySQL <span>[github.com/example/shorturl](https://github.com/example/shorturl)</span>

- 用发号器 + Base62 编码生成短链，单机压测 QPS 1.2 万，P99 < 15ms。
- 实现布隆过滤器拦截无效查询，缓存穿透场景下数据库压力下降 94%。
- 完整的 Docker Compose 部署与 Prometheus 监控面板，README 含压测报告。

### 校园二手书交易平台 <span>2024.10 - 2024.12</span>
#### 团队 4 人 · 担任后端负责人 <span>Spring Boot + Vue</span>

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
