<style>
/* ===========================================================
   Article — LaTeX \documentclass{article} 的默认味道
   参考：Computer Modern / Latin Modern 的字形与竖向节奏

   \maketitle 居中，\section 只加粗不画线，正文两端对齐，
   全靠留白撑层级。十套里最端庄的一套，
   适合学术、咨询、法律、国企这类看重「正式」的场合。
   =========================================================== */

:root {
    --font-body: var(--font-serif);

    --fs: 10.6pt;
    --lh: 1.44;
    --fs-name: 19pt;
    --fs-sec: 12.5pt;
    --fs-sub: calc(var(--fs) * 0.95);

    --page-x: 22mm;
    --page-y: 20mm;
    --sp-sec: 6.4mm;
    --sp-item: 3.6mm;
    --sp-line: 1.6mm;

    --muted: #454545;
    --rule: #000;
    --tint: #f2f2f0;
}

/* --- \maketitle --- */
h1 {
    text-align: center;
    font-weight: 400;
    font-size: var(--fs-name);
    letter-spacing: 0.06em;
    margin: 1.5mm 0 2.4mm;
}

h1 + p {
    text-align: center;
    color: var(--ink);
    font-size: calc(var(--fs) * 0.94);
    line-height: 1.5;
    margin: 0 0 var(--sp-sec);
    padding-bottom: 4.2mm;
    border-bottom: 0.4pt solid var(--rule);
}

/* --- \section：齐左、加粗、不画线 --- */
h2 {
    font-weight: 700;
    letter-spacing: 0.008em;
    margin: var(--sp-sec) 0 1.8mm;
}

/* --- \subsection / \subsubsection --- */
h3 {
    font-weight: 700;
    font-size: calc(var(--fs) * 1.02);
}

h4 {
    font-style: italic;
    color: var(--ink);
}

[data-split] > .x {
    font-style: normal;
    color: var(--muted);
}

/* article 的正文是两端对齐；简历里保留这一条，去掉首行缩进 */
p, li {
    text-align: justify;
    text-align-last: left;
    hyphens: auto;
}

/* \begin{quote}：两侧内缩、不画竖线 */
blockquote {
    margin: 2.4mm 6mm;
    padding: 0;
    border: 0;
    font-size: calc(var(--fs) * 0.96);
    color: var(--ink);
}

/* \itemize 的层级：• 然后 – */
ul { padding-left: 5.4mm; }
li::marker { color: var(--ink); }
li > ul { list-style-type: "–  "; }

/* \texttt{}：等宽，不做成标签底色 */
code {
    font-family: var(--font-mono);
    font-size: 0.86em;
    background: none;
    padding: 0;
}

/* \begin{tabular}：上下双横线，中规中矩 */
table {
    border-top: 0.4pt solid var(--rule);
    border-bottom: 0.4pt solid var(--rule);
    padding: 0.6mm 0;
}

td { padding: 0.9mm 0; }
td:first-child { font-weight: 400; font-style: italic; width: 24mm; }
</style>

# 顾清和

管理咨询顾问 ⋄ 上海 ⋄ 137-0000-0000 ⋄ guqinghe@example.com

## 概述

四年战略与运营咨询经验，覆盖消费品、医疗器械与工业制造。习惯把模糊的商业问题拆成可验证的假设，再用数据和一线访谈把结论钉死。

## 工作经历

### 某国际管理咨询公司 <span>2022.08 - 至今</span>
#### 咨询顾问（Consultant） <span>上海</span>

- **问题**：某乳制品集团连续三年增长停滞，管理层认为是渠道问题。
  **做法**：拆解 27 万条 POS 数据与 60 场终端访谈，发现真正的流失在于新品与主力品的价格带重叠。
  **影响**：重划四条价格带并砍掉 18 个 SKU，次年同店销售额提升 11.4%，毛利率提升 2.6 个百分点。
- **问题**：某医疗器械公司出海东南亚，三个市场同时铺开但均未盈利。
  **做法**：建立市场吸引力与进入难度双维模型，量化监管周期、经销商集中度与支付能力。
  **影响**：收缩至单一市场重投，18 个月内该市场转为正现金流，累计节省投入约 3,200 万元。
- 带 2 名分析师完成尽调支持，覆盖 5 个标的的商业尽职调查，其中 2 单最终成交。

### 某国际管理咨询公司 <span>2021.07 - 2022.07</span>
#### 分析师（Analyst） <span>上海</span>

- 独立负责三个项目的数据工作流，把重复性的数据清洗从人均每周 12 小时压缩到 2 小时。
- 主导某工业客户的成本基线测算，识别出 7 类可优化成本项，年化节约 8,700 万元。

## 实习经历

### 某投资银行 · 消费组 <span>2020.06 - 2020.09</span>
#### 暑期实习生 <span>北京</span>

- 参与两个 IPO 项目的行业章节撰写与可比公司分析。

## 教育背景

### 复旦大学 <span>2017.09 - 2021.06</span>
#### 经济学 · 本科 <span>GPA 3.85/4.0，排名 6/162</span>

交换：University of Manchester，商学院，2019 年秋季学期

## 技能与其他

| 方法 | 假设树、MECE、五力模型、市场规模测算、商业尽调 |
|:--|:--|
| 工具 | Excel 建模、SQL、Python（pandas）、Tableau、PowerPoint |
| 语言 | 中文（母语）、英语（雅思 7.5，工作语言） |
| 其他 | CFA 一级通过；曾任校辩论队队长 |
