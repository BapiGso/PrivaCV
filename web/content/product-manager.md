<style>
/* ===========================================================
   Awesome CV — 复刻 github.com/posquit0/Awesome-CV
   GitHub 上 star 最多的那份 LaTeX 简历。

   签名细节：巨大的居中姓名（名细体、姓粗体），章节标题着色、
   后面拖一条横贯到页边的细线。全篇只有一个强调色。

   · 双色姓名写成 `# Wei **Chen**`
   · 联系方式行里用 *斜体* 标出的部分会变成强调色（放职位方向）
   · 章节标题请保持纯文本，别在 ## 里加粗体或链接
   =========================================================== */

:root {
    --accent: #b23a2e;                 /* 原版的 awesome-red，换成 #16697a 就是 awesome-emerald */
    --tint: rgba(178, 58, 46, 0.09);   /* 换 --accent 时记得跟着换 */
}

body {
    box-sizing: border-box;
    padding: 15mm 18mm;
    font-family: Inter, "Source Sans 3", "Helvetica Neue", Helvetica, Arial,
                 "Segoe UI", Roboto, "PingFang SC", "Source Han Sans SC",
                 "Microsoft YaHei", sans-serif;
    font-size: 9.9pt;
    line-height: 1.5;
    color: #2b2b2b;
}

h1 {
    text-align: center;
    font-size: 27pt;
    font-weight: 300;
    line-height: 1.15;
    letter-spacing: 0.015em;
    margin: 1mm 0 2mm;
}

h1 strong { font-weight: 700; }

h1 + p {
    text-align: center;
    font-size: 8.91pt;
    line-height: 1.6;
    color: #6f6f6f;
    margin: 0 0 5.6mm;
}

/* 第一行写成 *高级后端工程师 ⋄ 分布式系统* 就是原版那条彩色定位语 */
h1 + p em {
    font-style: normal;
    color: var(--accent);
    letter-spacing: 0.09em;
    text-transform: uppercase;
    font-size: 0.94em;
}

/* --- 着色章节 + 右侧延长线 --- */
h2 {
    display: flex;
    align-items: center;
    gap: 3.5mm;
    color: var(--accent);
    font-size: 12.5pt;
    font-weight: 700;
    line-height: 1.25;
    text-transform: uppercase;
    letter-spacing: 0.07em;
    margin: 5.6mm 0 2mm;
}

h2::after {
    content: "";
    flex: 1 1 auto;
    border-top: 0.8pt solid #d8d8d8;
}

h2:first-of-type { margin-top: 0; }

h3 {
    font-size: 9.9pt;
    font-weight: 600;
    line-height: 1.3;
    color: #111;
    margin: 3.4mm 0 0;
}

h4 {
    font-size: 9.11pt;
    font-weight: 400;
    font-style: italic;
    color: #6f6f6f;
    line-height: 1.35;
    margin: 0.4mm 0 1.4mm;
}

/* 时间和地点：斜体写在标题末尾，浮到右边 */
h3 em, h4 em {
    float: right;
    font-weight: 400;
    font-size: 0.95em;
    color: #6f6f6f;
    white-space: nowrap;
    font-variant-numeric: lining-nums tabular-nums;
}

h3 em {
    font-style: normal;
    font-weight: 600;
    color: var(--accent);
}

p { margin: 0 0 1.4mm; }
ul, ol { margin: 1.4mm 0 2mm; padding-left: 4.6mm; }
li { margin: 0 0 0.9mm; }
li::marker { color: var(--accent); font-size: 0.9em; }
hr { border: 0; border-top: 0.5pt solid #d8d8d8; margin: 3mm 0; }

a { color: var(--accent); text-decoration: none; }
h1 + p a { color: inherit; }

/* 行内 code 当技能标签用：浅底圆角，字体跟正文走 */
code {
    font-family: inherit;
    font-size: 0.92em;
    padding: 0.4mm 1.6mm;
    border-radius: 1.4mm;
    background: var(--tint);
    white-space: nowrap;
    color: var(--accent);
}

blockquote {
    margin: 1.4mm 0;
    padding: 0 0 0 3.5mm;
    border-left: 1.2pt solid var(--accent);
    color: #6f6f6f;
}

blockquote p:last-child { margin-bottom: 0; }

/* GFM 要求表格有表头行，留空即可；这里把它藏起来 */
table {
    width: 100%;
    border-collapse: collapse;
    font-size: 0.96em;
    margin: 1.4mm 0 2.2mm;
}

thead { display: none; }
td { padding: 0.7mm 0; vertical-align: top; }

td:first-child {
    width: 1px;
    white-space: nowrap;
    padding-right: 4mm;
    color: var(--accent);
    font-weight: 600;
}
</style>

# 苏文彦

产品经理 · 5 年经验 ⋄ 杭州 ⋄ 136-0000-0000 ⋄ suwenyan@example.com

> B 端 SaaS 方向，擅长从混乱的业务流程里抽出可复用的产品结构；做过 0→1，也做过把一个亏损产品线拉回正毛利。

## 工作经历

### 某企业服务公司 *2022.03 - 至今*
#### 高级产品经理 · 供应链协同产品线 *杭州*

- 负责供应商协同平台从 0 到 1：8 个月内完成需求调研、MVP 上线到首批 30 家客户签约，首年 ARR 620 万元。
- 通过 26 场客户走访重新定义核心场景，砍掉原规划中 40% 的功能，把首版交付周期从 11 个月压缩到 8 个月。
- 设计对账自动化模块，客户平均对账工时从每月 40 小时降到 6 小时，成为续约率最高的功能模块（续约提及率 78%）。
- 主导定价模型改版，从「按账号数」改为「按单据量 + 基础包」，客单价提升 34%，小客户流失率反而下降 9%。

### 某物流科技公司 *2020.05 - 2022.02*
#### 产品经理 *杭州*

- 负责运单管理后台，梳理 7 个历史遗留状态机为 1 套统一模型，客服因状态混乱产生的工单下降 62%。
- 推动移动端司机 App 的接单流程改版，从 6 步减到 3 步，接单转化率提升 21%。
- 与算法团队共建智能调度功能，定义了可解释的调度理由展示，调度员采纳率从 43% 提升到 81%。

## 项目经历

### 客户健康度体系 *2023*
#### 主导 · 跨部门项目 *产品 + 数据 + CS*

- 联合数据团队定义健康度分模型（使用深度、活跃席位、工单情绪三个维度），提前 60 天识别流失风险客户。
- CS 团队据此做主动干预，试点季度的续约率相比对照组提升 14 个百分点。

## 教育背景

### 浙江大学 *2016.09 - 2020.06*
#### 工业工程 · 本科 *GPA 3.7/4.0*

## 技能与其他

|  |  |
|:--|:--|
| 方法 | 用户访谈、Jobs-to-be-Done、服务蓝图、需求优先级（RICE / Kano） |
| 工具 | Figma、Axure、SQL（能自己跑数）、Tableau、Jira |
| 认证 | NPDP 产品经理国际资格认证 |
| 语言 | 英语 CET-6，可英文邮件与文档往来 |
