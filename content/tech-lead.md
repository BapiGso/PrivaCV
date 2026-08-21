<style>
/* ===========================================================
   Classic — 复刻 moderncv 的 classic 样式
   （\moderncvstyle{classic}）

   签名细节：时间戳单独占左边一列，正文全部对齐到那一列右边，
   章节标题反过来悬在左页边上。这是最有「LaTeX 版式」感的一套，
   多份经历排下来时间线一眼看得清。

   注意：条目请写成 `### 公司 <span>2021.06 - 至今</span>`，
   没有 <span> 的三级标题会自动缩到正文列，不会歪。
   =========================================================== */

:root {
    --font-body: var(--font-serif);
    --date-col: 27mm;

    --accent: #6b4f2a;
    --fs: 10.3pt;
    --lh: 1.45;
    --fs-name: 20pt;
    --fs-sec: 11pt;

    --page-x: 16mm;
    --page-y: 16mm;
    --sp-sec: 5.6mm;
    --sp-item: 3.4mm;

    --ink: #1c1c1c;
    --muted: #5b5b5b;
    --rule-soft: #dcd6cc;
    --tint: rgba(107, 79, 42, 0.09);
}

/* 正文整体让出左边的时间列；标题和头部不让 */
body > * { margin-left: var(--date-col); }

body > h1,
body > h1 + p,
body > h2,
body > h3[data-split],
body > hr,
body > .page-break { margin-left: 0; }

h1 {
    font-weight: 400;
    letter-spacing: 0.03em;
    margin: 0 0 1.6mm;
}

h1 + p {
    color: var(--muted);
    font-size: calc(var(--fs) * 0.92);
    line-height: 1.5;
    margin: 0 0 var(--sp-sec);
    padding-bottom: 3mm;
    border-bottom: 0.5pt solid var(--rule-soft);
}

/* 章节标题悬在左页边，正好占满时间列 */
h2 {
    color: var(--accent);
    font-weight: 700;
    font-variant: small-caps;
    letter-spacing: 0.06em;
    margin: var(--sp-sec) 0 2mm;
}

/* --- 条目：时间挪到左列 --- */
h3 { font-weight: 700; }

h3[data-split] { gap: 0; }

h3[data-split] > .x {
    order: -1;
    flex: 0 0 var(--date-col);
    margin-left: 0;
    padding-right: 4mm;
    text-align: left;
    color: var(--accent);
    font-size: 0.92em;
}

h4 {
    font-style: italic;
    color: var(--muted);
}

li::marker { color: var(--accent); }

a { color: var(--accent); }
h1 + p a { color: inherit; }

code {
    font-family: var(--font-body);
    color: var(--accent);
    background: none;
    padding: 0;
    font-style: italic;
}

blockquote {
    border-left-color: var(--accent);
    font-style: italic;
}

td:first-child {
    font-weight: 400;
    font-style: italic;
    color: var(--accent);
}
</style>

# 陈牧之

技术负责人 · 10 年经验 ⋄ 深圳 ⋄ 139-0000-0000 ⋄ chenmuzhi@example.com ⋄ [github.com/example](https://github.com/example)

## 概述

从一线工程师做到 30 人研发团队负责人。擅长在业务快速扩张期把技术债、交付节奏和人员梯队三件事同时管住，也愿意继续写关键路径上的代码。

## 工作经历

### 某跨境电商 · 技术中心 <span>2021.02 - 至今</span>
#### 研发总监 · 交易与履约 <span>深圳</span>

- 接手时 3 个团队 18 人、需求交付周期中位数 34 天；重建需求分级与双周节奏后压到 11 天，同期人员规模翻倍到 34 人。
- 主导交易系统重构，把 7 个耦合服务收敛为 3 个领域服务，大促期间订单峰值从 3 万 QPS 撑到 9 万，可用性 99.99%。
- 建立技术雷达与季度架构评审，两年内下线 12 个僵尸服务，云成本年化下降 480 万元。
- 搭建工程师职级与晋升标准，两年培养出 4 名 team leader，主动离职率从 21% 降到 7%。

### 某本地生活平台 <span>2018.05 - 2021.01</span>
#### 技术经理 · 商家平台 <span>广州</span>

- 带 9 人团队从零搭建商家开放平台，一年内接入 ISV 200+，平台调用量日均 4,000 万次。
- 推动全组接入契约测试与灰度发布，线上回滚次数从每月 6 次降到 0.5 次。
- 牵头一次严重线上故障的复盘，落地容量水位告警与降级预案，同类故障此后未再发生。

### 某支付公司 <span>2015.07 - 2018.04</span>
#### 高级后端工程师 → 技术专家 <span>广州</span>

- 负责清结算核心链路，设计对账差错自动追平机制，人工介入率从 3.2% 降到 0.04%。
- 主导数据库分库分表迁移，8 亿存量数据零停机切换。

## 项目经历

### 研发效能度量平台 <span>2022 - 2023</span>
#### 发起人 · 跨团队 <span>内部开源</span>

- 用 DORA 四项指标替代原来的「代码行数 / 工时」考核，推广到 6 个部门 120 人。
- 度量口径公开可查，避免了指标博弈；部署频率提升 3.1 倍，变更失败率下降 42%。

## 教育背景

### 中山大学 <span>2011.09 - 2015.06</span>
#### 软件工程 · 本科 <span>GPA 3.6/4.0</span>

## 技能与其他

| 技术 | 分布式系统、服务治理、数据库、可观测性、成本治理 |
|:--|:--|
| 管理 | 目标拆解、职级体系、绩效校准、跨部门协作、招聘面试 |
| 语言 | 中文（母语）、英语（可主持英文会议） |
| 输出 | ArchSummit 2023 讲师；内部技术分享 40+ 场 |
