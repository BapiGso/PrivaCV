<style>
/* ===========================================================
   Article — LaTeX \documentclass{article} 的默认味道
   参考：Computer Modern / Latin Modern 的字形与竖向节奏

   \maketitle 居中，\section 只加粗不画线，正文两端对齐，
   全靠留白撑层级。十套里最端庄的一套，
   适合学术、咨询、法律、国企这类看重「正式」的场合。
   =========================================================== */

body {
    padding: 20mm 22mm;
    font-family: "Latin Modern Roman", XCharter, Charter, "Bitstream Charter",
                 Palatino, Georgia, "Times New Roman", "Source Han Serif SC",
                 "Songti SC", serif;
    font-size: 10.6pt;
    line-height: 1.44;
    color: #111;
}

/* --- \maketitle --- */
h1 {
    text-align: center;
    font-weight: 400;
    font-size: 19pt;
    letter-spacing: 0.06em;
    line-height: 1.15;
    margin: 1.5mm 0 2.4mm;
}

h1 + p {
    text-align: center;
    color: #111;
    font-size: 9.96pt;
    line-height: 1.5;
    margin: 0 0 6.4mm;
    padding-bottom: 4.2mm;
    border-bottom: 0.4pt solid #000;
}

/* --- \section：齐左、加粗、不画线 --- */
h2 {
    font-weight: 700;
    font-size: 12.5pt;
    line-height: 1.25;
    letter-spacing: 0.008em;
    margin: 6.4mm 0 1.8mm;
}

h2:first-of-type { margin-top: 0; }

/* --- \subsection / \subsubsection --- */
h3 {
    font-weight: 700;
    font-size: 10.81pt;
    line-height: 1.3;
    margin: 3.6mm 0 0;
}

h4 {
    font-style: italic;
    font-weight: 400;
    font-size: 10.07pt;
    line-height: 1.35;
    color: #111;
    margin: 0.4mm 0 1.6mm;
}

/* 时间和地点：斜体写在标题末尾，浮到右边 */
h3 em, h4 em {
    float: right;
    font-style: normal;
    font-weight: 400;
    font-variant-numeric: tabular-nums;
    color: #454545;
}

/* article 的正文是两端对齐；简历里保留这一条，去掉首行缩进 */
p, li {
    text-align: justify;
    text-align-last: left;
    hyphens: auto;
}

p { margin: 0 0 1.6mm; }

/* \begin{quote}：两侧内缩、不画竖线 */
blockquote {
    margin: 2.4mm 6mm;
    padding: 0;
    border: 0;
    font-size: 10.18pt;
    color: #111;
}

/* \itemize 的层级：• 然后 – */
ul, ol { margin: 1.6mm 0 2.2mm; padding-left: 4.6mm; }
ul { padding-left: 5.4mm; }
li { margin: 0 0 0.9mm; }
li::marker { color: #111; font-size: 0.9em; }
li > ul { list-style-type: "–  "; }

a { color: inherit; text-decoration: none; }

/* \texttt{}：等宽，不做成标签底色 */
code {
    font-family: inherit;
    font-size: 0.86em;
    padding: 0.3mm 1.4mm;
    border-radius: 0.8mm;
    background: rgba(17, 17, 17, 0.06);
    white-space: nowrap;
}

/* \begin{tabular}：上下双横线，中规中矩 */
table {
    width: 100%;
    border-collapse: collapse;
    font-size: 0.96em;
    margin: 1.6mm 0 2.4mm;
    border-top: 0.4pt solid #000;
    border-bottom: 0.4pt solid #000;
    padding: 0.6mm 0;
}

thead { display: none; }
td { padding: 0.9mm 0; vertical-align: top; }

td:first-child {
    width: 1px;
    white-space: nowrap;
    padding-right: 5mm;
    font-weight: 400;
    font-style: italic;
}
</style>

# 顾清和

管理咨询顾问 ⋄ 上海 ⋄ 137-0000-0000 ⋄ guqinghe@example.com

## 概述

四年战略与运营咨询经验，覆盖消费品、医疗器械与工业制造。习惯把模糊的商业问题拆成可验证的假设，再用数据和一线访谈把结论钉死。

## 工作经历

### 某国际管理咨询公司 *2022.08 - 至今*
#### 咨询顾问（Consultant） *上海*

- **问题**：某乳制品集团连续三年增长停滞，管理层认为是渠道问题。
  **做法**：拆解 27 万条 POS 数据与 60 场终端访谈，发现真正的流失在于新品与主力品的价格带重叠。
  **影响**：重划四条价格带并砍掉 18 个 SKU，次年同店销售额提升 11.4%，毛利率提升 2.6 个百分点。
- **问题**：某医疗器械公司出海东南亚，三个市场同时铺开但均未盈利。
  **做法**：建立市场吸引力与进入难度双维模型，量化监管周期、经销商集中度与支付能力。
  **影响**：收缩至单一市场重投，18 个月内该市场转为正现金流，累计节省投入约 3,200 万元。
- 带 2 名分析师完成尽调支持，覆盖 5 个标的的商业尽职调查，其中 2 单最终成交。

### 某国际管理咨询公司 *2021.07 - 2022.07*
#### 分析师（Analyst） *上海*

- 独立负责三个项目的数据工作流，把重复性的数据清洗从人均每周 12 小时压缩到 2 小时。
- 主导某工业客户的成本基线测算，识别出 7 类可优化成本项，年化节约 8,700 万元。

## 实习经历

### 某投资银行 · 消费组 *2020.06 - 2020.09*
#### 暑期实习生 *北京*

- 参与两个 IPO 项目的行业章节撰写与可比公司分析。

## 教育背景

### 复旦大学 *2017.09 - 2021.06*
#### 经济学 · 本科 *GPA 3.85/4.0，排名 6/162*

交换：University of Manchester，商学院，2019 年秋季学期

## 技能与其他

|  |  |
|:--|:--|
| 方法 | 假设树、MECE、五力模型、市场规模测算、商业尽调 |
| 工具 | Excel 建模、SQL、Python（pandas）、Tableau、PowerPoint |
| 语言 | 中文（母语）、英语（雅思 7.5，工作语言） |
| 其他 | CFA 一级通过；曾任校辩论队队长 |
