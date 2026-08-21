<style>
/* ===========================================================
   Jake's Resume — 复刻 github.com/jakeryang/resume
   （Overleaf 上下载量最高的那份，MIT）

   签名细节：居中的小型大写姓名、\scshape 章节标题下面一条通栏
   \titlerule、条目「粗体标题 · 右侧时间」＋「斜体职位 · 右侧地点」。
   纯黑白单栏，ATS 解析友好，工程岗海投的默认选择。
   =========================================================== */

:root {
    --font-body: var(--font-serif);

    --fs: 10.4pt;
    --lh: 1.38;
    --fs-name: 23pt;
    --fs-sec: 12pt;

    --page-x: 15mm;
    --page-y: 13mm;
    --sp-sec: 4.8mm;
    --sp-item: 2.8mm;
    --sp-line: 1mm;

    --rule: #000;
    --muted: #1a1a1a;
    --tint: #f0f0ee;
}

/* --- \Huge\scshape 居中姓名 --- */
h1 {
    text-align: center;
    font-variant: small-caps;
    font-weight: 400;
    letter-spacing: 0.035em;
    margin: 0 0 1.4mm;
}

h1 + p {
    text-align: center;
    color: var(--ink);
    font-size: calc(var(--fs) * 0.92);
    line-height: 1.45;
    margin: 0 0 var(--sp-sec);
}

/* 原版把链接加了下划线，这里保留 */
h1 + p a {
    text-decoration: underline;
    text-underline-offset: 1.6pt;
    text-decoration-thickness: 0.4pt;
}

/* --- \section + \titlerule --- */
h2 {
    font-variant: small-caps;
    font-weight: 400;
    letter-spacing: 0.04em;
    padding-bottom: 0.5mm;
    border-bottom: 0.6pt solid var(--rule);
    margin: var(--sp-sec) 0 1.4mm;
}

/* --- \resumeSubheading --- */
h3 { font-weight: 700; }

h4 {
    font-style: italic;
    color: var(--ink);
    font-size: calc(var(--fs) * 0.94);
}

h3 > .x { color: var(--ink); }
h4 > .x { font-style: italic; color: var(--ink); }

/* 原版的 item 是个很小的 \bullet，缩进也很浅 */
ul { padding-left: 5.2mm; }
li::marker { font-size: 0.7em; color: var(--ink); }
li { margin-bottom: 0.7mm; }

code {
    font-family: var(--font-body);
    background: none;
    padding: 0;
    font-size: 1em;
}

td:first-child { width: 22mm; font-weight: 700; }
</style>

# 张明远

后端工程师 · 6 年经验 ⋄ 上海 ⋄ 138-0000-0000 ⋄ zhangmingyuan@example.com ⋄ [github.com/example](https://github.com/example)

## 技能

- 语言：`Go` `Java` `Python` `SQL`
- 存储：`MySQL` `Redis` `Kafka` `Elasticsearch` `ClickHouse`
- 基建：`Kubernetes` `Docker` `Prometheus` `Grafana` `ArgoCD`
- 方向：分布式系统、高并发、服务治理、可观测性

## 工作经历

### 某电商平台 · 基础架构部 <span>2021.06 - 至今</span>
#### 高级后端工程师 <span>上海</span>

- 主导订单中心从单体拆分为 6 个微服务，梳理出 40+ 处隐式耦合，迁移期间零事故，发布频率从每两周一次提升到每天 3 次。
- 重构订单查询链路：引入 ClickHouse 做冷数据归档 + Redis 二级缓存，P99 从 1.8s 降到 210ms，MySQL 主库 QPS 下降 63%。
- 设计并落地全链路压测方案，在大促前发现 3 个会导致雪崩的线程池配置问题；当年双十一峰值 8.2 万 QPS，服务可用性 99.99%。
- 推动组内接入 OpenTelemetry，把平均故障定位时间（MTTR）从 47 分钟压到 12 分钟。

### 某出行科技公司 <span>2019.03 - 2021.05</span>
#### 后端工程师 <span>杭州</span>

- 负责派单引擎的匹配模块，用空间索引替换全量扫描，单次匹配耗时从 380ms 降到 45ms，支撑日均 200 万订单。
- 搭建基于 Kafka 的订单事件总线，替换原有的定时轮询同步，下游 5 个业务方的数据延迟从分钟级降到秒级。
- 主动排查并修复一个长期存在的分布式锁误释放问题，该问题此前每月造成约 300 单重复派送。

## 项目经历

### 轻量级配置中心（开源，1.2k Star） <span>2022 - 至今</span>
#### 个人项目 · Go <span>[github.com/example/conf](https://github.com/example/conf)</span>

- 基于 Raft 实现的配置中心，支持灰度发布与秒级回滚，单节点内存占用 < 30MB。
- 被 3 家公司用于生产环境，累计处理 issue 80+，合并外部 PR 24 个。

## 教育背景

### 华东师范大学 <span>2015.09 - 2019.06</span>
#### 计算机科学与技术 · 本科 <span>GPA 3.7/4.0</span>

## 其他

| 标签 | 内容 |
|:--|:--|
| 语言 | 中文（母语）、英语（CET-6，可无障碍阅读技术文档） |
| 输出 | 技术博客 60+ 篇，累计阅读 30 万；GopherChina 2023 分享嘉宾 |
