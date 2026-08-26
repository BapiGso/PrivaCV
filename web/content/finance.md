<style>
/* ===========================================================
   sb2nov — 复刻 github.com/sb2nov/resume
   北美技术圈的另一份「事实标准」，比 Jake 更靠左、更书面。

   签名细节：姓名齐左不居中，头部下面一对 \hrule（CSS 里就是
   double 边框），章节标题全大写 + 字距拉开。
   =========================================================== */

body {
    padding: 15mm 17mm;
    font-family: "Latin Modern Roman", XCharter, Charter, "Bitstream Charter",
                 Palatino, Georgia, "Times New Roman", "Source Han Serif SC",
                 "Songti SC", serif;
    font-size: 10.5pt;
    line-height: 1.42;
    color: #111;
}

h1 {
    font-weight: 700;
    font-size: 20pt;
    letter-spacing: 0.01em;
    line-height: 1.15;
    margin: 0 0 1.2mm;
}

/* LaTeX 里常写成叠一对 \hrule，CSS 的 double 边框正好是这个效果 */
h1 + p {
    color: #3d3d3d;
    font-size: 9.66pt;
    line-height: 1.5;
    margin: 0 0 5.2mm;
    padding-bottom: 2.6mm;
    border-bottom: 3.5pt double #000;
}

h2 {
    font-weight: 700;
    text-transform: uppercase;
    letter-spacing: 0.14em;
    font-size: 11pt;
    line-height: 1.25;
    padding-bottom: 0.7mm;
    border-bottom: 0.5pt solid #000;
    margin: 5.2mm 0 1.6mm;
}

h2:first-of-type { margin-top: 0; }

h3 {
    font-weight: 700;
    font-size: 10.5pt;
    line-height: 1.3;
    margin: 3mm 0 0;
}

h4 {
    font-style: italic;
    font-weight: 400;
    font-size: 9.66pt;
    line-height: 1.35;
    color: #3d3d3d;
    margin: 0.4mm 0 1.2mm;
}

/* 时间和地点：斜体写在标题末尾，浮到右边 */
h3 em, h4 em { float: right; font-weight: 400; font-variant-numeric: tabular-nums; }
h3 em { font-style: italic; color: #111; }

p { margin: 0 0 1.2mm; }
ul, ol { margin: 1.2mm 0 1.8mm; padding-left: 4.6mm; }
ul { padding-left: 5mm; }
li { margin: 0 0 0.9mm; }
li::marker { font-size: 0.75em; color: #111; }

a { color: inherit; text-decoration: none; }

code {
    font-family: inherit;
    font-style: italic;
    font-size: 0.92em;
    padding: 0.3mm 1.4mm;
    border-radius: 0.8mm;
    background: rgba(17, 17, 17, 0.06);
    white-space: nowrap;
}

table { width: 100%; border-collapse: collapse; font-size: 0.96em; margin: 1.2mm 0 2mm; }
thead { display: none; }
td { padding: 0.7mm 0; vertical-align: top; }

td:first-child {
    width: 1px;
    white-space: nowrap;
    padding-right: 5mm;
    padding-top: 1.1mm;
    font-weight: 400;
    text-transform: uppercase;
    letter-spacing: 0.08em;
    font-size: 0.88em;
    color: #3d3d3d;
}
</style>

# 顾清源

财务分析 / 投资分析 ⋄ 上海 ⋄ 135-0000-0000 ⋄ guqingyuan@example.com ⋄ CPA（已通过全科）

## 教育背景

### 上海财经大学 *2017.09 - 2020.06*
#### 金融学 · 硕士 *GPA 3.8/4.0*

### 厦门大学 *2013.09 - 2017.06*
#### 会计学 · 本科 *专业排名 3/126*

## 工作经历

### 某消费行业上市公司 · 财务部 *2022.04 - 至今*
#### 高级财务分析师 *上海*

- 负责集团年度预算与滚动预测，重构预测模型的驱动逻辑（从科目外推改为量价驱动），全年收入预测偏差从 ±11% 收窄到 ±3.5%。
- 主导 3 条产品线的盈利能力分析，识别出一条长期被共同成本掩盖的负毛利业务，推动关停后年化改善经营利润约 2,300 万元。
- 搭建管理报表自动化流程（SAP 取数 + Power Query + Power BI），月结报表出具时间从 T+8 缩短到 T+3。
- 参与一起 4.5 亿元的同业并购，负责标的财务尽调与估值建模（DCF / 可比公司），发现两项未披露的关联方往来。

### 某会计师事务所（内资八大） *2020.07 - 2022.03*
#### 审计员 → 高级审计员 *上海*

- 独立负责 6 家制造业与消费业客户的年报审计现场，累计带 3 人小组。
- 在一家客户的存货监盘中发现跨期确认收入 1,800 万元，出具审计调整并促成客户修订内控流程。

## 技能

|  |  |
|:--|:--|
| 专业 | 财务建模、估值（DCF / 可比公司 / 先例交易）、预算与滚动预测、并购尽调 |
| 准则 | 中国企业会计准则、IFRS |
| 工具 | Excel（数组公式 / Power Query / VBA）、SAP FICO、Power BI、Wind、Capital IQ |
| 证书 | CPA（全科通过）、CFA 二级 |
| 语言 | 英语 CET-6 / TOEFL 105，可英文汇报与阅读英文年报 |

## 其他

- 校招期间获全国大学生会计信息化技能大赛一等奖
- 业余维护一个财报拆解的公众号，累计 120 篇，订阅 1.5 万
