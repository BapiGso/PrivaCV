<style>
/* ===========================================================
   sb2nov — 复刻 github.com/sb2nov/resume
   北美技术圈的另一份「事实标准」，比 Jake 更靠左、更书面。

   签名细节：姓名齐左不居中，头部下面一对 \hrule（CSS 里就是
   double 边框），章节标题全大写 + 字距拉开。
   =========================================================== */

:root {
    --font-body: var(--font-serif);

    --fs: 10.5pt;
    --lh: 1.42;
    --fs-name: 20pt;
    --fs-sec: 11pt;

    --page-x: 17mm;
    --page-y: 15mm;
    --sp-sec: 5.2mm;
    --sp-item: 3mm;
    --sp-line: 1.2mm;

    --rule: #000;
    --muted: #3d3d3d;
    --tint: #f1f1ef;
}

h1 {
    font-weight: 700;
    letter-spacing: 0.01em;
    margin: 0 0 1.2mm;
}

/* LaTeX 里常写成叠一对 \hrule，CSS 的 double 边框正好是这个效果 */
h1 + p {
    color: var(--muted);
    font-size: calc(var(--fs) * 0.92);
    line-height: 1.5;
    margin: 0 0 var(--sp-sec);
    padding-bottom: 2.6mm;
    border-bottom: 3.5pt double var(--rule);
}

h2 {
    font-weight: 700;
    text-transform: uppercase;
    letter-spacing: 0.14em;
    font-size: var(--fs-sec);
    padding-bottom: 0.7mm;
    border-bottom: 0.5pt solid var(--rule);
    margin: var(--sp-sec) 0 1.6mm;
}

h3 { font-weight: 700; }

h4 {
    font-style: italic;
    color: var(--muted);
}

h3 > .x {
    font-style: italic;
    color: var(--ink);
}

ul { padding-left: 5mm; }
li::marker { font-size: 0.75em; color: var(--ink); }

code {
    font-family: var(--font-body);
    background: none;
    padding: 0;
    font-style: italic;
}

td:first-child {
    width: 24mm;
    font-weight: 400;
    text-transform: uppercase;
    letter-spacing: 0.08em;
    font-size: 0.88em;
    color: var(--muted);
    padding-top: 1.1mm;
}
</style>

# 顾清源

财务分析 / 投资分析 ⋄ 上海 ⋄ 135-0000-0000 ⋄ guqingyuan@example.com ⋄ CPA（已通过全科）

## 教育背景

### 上海财经大学 <span>2017.09 - 2020.06</span>
#### 金融学 · 硕士 <span>GPA 3.8/4.0</span>

### 厦门大学 <span>2013.09 - 2017.06</span>
#### 会计学 · 本科 <span>专业排名 3/126</span>

## 工作经历

### 某消费行业上市公司 · 财务部 <span>2022.04 - 至今</span>
#### 高级财务分析师 <span>上海</span>

- 负责集团年度预算与滚动预测，重构预测模型的驱动逻辑（从科目外推改为量价驱动），全年收入预测偏差从 ±11% 收窄到 ±3.5%。
- 主导 3 条产品线的盈利能力分析，识别出一条长期被共同成本掩盖的负毛利业务，推动关停后年化改善经营利润约 2,300 万元。
- 搭建管理报表自动化流程（SAP 取数 + Power Query + Power BI），月结报表出具时间从 T+8 缩短到 T+3。
- 参与一起 4.5 亿元的同业并购，负责标的财务尽调与估值建模（DCF / 可比公司），发现两项未披露的关联方往来。

### 某会计师事务所（内资八大） <span>2020.07 - 2022.03</span>
#### 审计员 → 高级审计员 <span>上海</span>

- 独立负责 6 家制造业与消费业客户的年报审计现场，累计带 3 人小组。
- 在一家客户的存货监盘中发现跨期确认收入 1,800 万元，出具审计调整并促成客户修订内控流程。

## 技能

| 类别 | 内容 |
|:--|:--|
| 专业 | 财务建模、估值（DCF / 可比公司 / 先例交易）、预算与滚动预测、并购尽调 |
| 准则 | 中国企业会计准则、IFRS |
| 工具 | Excel（数组公式 / Power Query / VBA）、SAP FICO、Power BI、Wind、Capital IQ |
| 证书 | CPA（全科通过）、CFA 二级 |
| 语言 | 英语 CET-6 / TOEFL 105，可英文汇报与阅读英文年报 |

## 其他

- 校招期间获全国大学生会计信息化技能大赛一等奖
- 业余维护一个财报拆解的公众号，累计 120 篇，订阅 1.5 万
