<style>
/* ===========================================================
   Banking — 复刻 CTAN 上 moderncv 宏包的 banking 样式
   （\moderncvstyle{banking}\moderncvcolor{blue}）

   签名细节：姓名齐左、细体大字，头部下面一条粗强调色横线，
   章节标题着色全大写、下面一条发丝线。克制但一眼认得出。
   =========================================================== */

:root {
    --accent: #2b5c8a;                /* moderncv blue，可换 #7a2e2e / #2f6b46 */
    --tint: rgba(43, 92, 138, 0.09);  /* 换 --accent 时记得跟着换 */
}

body {
    box-sizing: border-box;
    padding: 15mm 17mm;
    font-family: Inter, "Source Sans 3", "Helvetica Neue", Helvetica, Arial,
                 "Segoe UI", Roboto, "PingFang SC", "Source Han Sans SC",
                 "Microsoft YaHei", sans-serif;
    font-size: 10.1pt;
    line-height: 1.48;
    color: #232323;
}

h1 {
    font-size: 24pt;
    font-weight: 300;
    line-height: 1.15;
    letter-spacing: -0.008em;
    margin: 0 0 1.6mm;
}

h1 strong {
    font-weight: 600;
    color: var(--accent);
}

h1 + p {
    font-size: 9.09pt;
    line-height: 1.55;
    color: #616161;
    margin: 0 0 5.4mm;
    padding-bottom: 3.2mm;
    border-bottom: 2.4pt solid var(--accent);
}

h2 {
    color: var(--accent);
    font-size: 10.8pt;
    font-weight: 700;
    line-height: 1.25;
    text-transform: uppercase;
    letter-spacing: 0.09em;
    padding-bottom: 1mm;
    border-bottom: 0.5pt solid #d5dde5;
    margin: 5.4mm 0 2mm;
}

h2:first-of-type { margin-top: 0; }

h3 {
    font-size: 10.1pt;
    font-weight: 600;
    line-height: 1.3;
    margin: 3.2mm 0 0;
}

h4 {
    font-size: 9.29pt;
    font-weight: 400;
    color: #616161;
    line-height: 1.35;
    margin: 0.4mm 0 1.4mm;
}

/* 时间和地点：斜体写在标题末尾，浮到右边 */
h3 em, h4 em {
    float: right;
    font-style: normal;
    font-weight: 400;
    font-size: 0.95em;
    color: #616161;
    white-space: nowrap;
    font-variant-numeric: lining-nums tabular-nums;
}

h3 em {
    color: var(--accent);
    font-weight: 500;
}

p { margin: 0 0 1.4mm; }
ul, ol { margin: 1.4mm 0 2mm; padding-left: 4.6mm; }
li { margin: 0 0 0.9mm; }
li::marker { color: var(--accent); font-size: 0.9em; }
hr { border: 0; border-top: 0.5pt solid #d5dde5; margin: 3mm 0; }

a { color: var(--accent); text-decoration: none; }
h1 + p a { color: inherit; }

/* 行内 code 当技能标签用：浅底圆角，字体跟正文走 */
code {
    font-family: inherit;
    font-size: 0.92em;
    padding: 0.3mm 1.4mm;
    border-radius: 0.8mm;
    background: var(--tint);
    white-space: nowrap;
    color: var(--accent);
}

blockquote {
    margin: 1.4mm 0;
    padding: 0 0 0 3.5mm;
    border-left: 1.2pt solid var(--accent);
    color: #616161;
    font-style: italic;
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
    font-weight: 600;
    color: var(--accent);
}
</style>

# 陈亦然

数据分析师 · 4 年经验 ⋄ 北京 ⋄ 137-0000-0000 ⋄ chenyiran@example.com

## 技能

|  |  |
|:--|:--|
| 分析 | SQL（复杂窗口函数、性能调优）、Python（pandas / statsmodels / scikit-learn） |
| 实验 | A/B 实验设计、假设检验、因果推断（DID、PSM）、样本量估算 |
| 可视化 | Tableau、Superset、Matplotlib、ECharts |
| 数据栈 | Hive、Spark、ClickHouse、dbt、Airflow |

## 工作经历

### 某内容社区 · 增长数据组 *2022.08 - 至今*
#### 数据分析师 *北京*

- 搭建新用户 7 日留存的归因分析框架，定位到「首次关注引导」是留存的关键路径；据此推动的产品改版使 D7 留存从 24.1% 提升到 29.6%。
- 重构核心指标口径与看板，把散落在 4 个团队的 30 余个「日活」定义收敛为 1 套，月度对账争议从平均 11 次降到 0。
- 负责社区激励预算的效果评估，用 DID 方法识别出 2 个 ROI 为负的补贴场景，全年节省预算约 480 万元。
- 建立 A/B 实验平台的元分析机制，复盘全年 213 场实验，发现 37% 的实验因样本量不足无法得出结论，推动实验准入规范落地。

### 某零售集团 *2021.06 - 2022.07*
#### 商业分析师 *北京*

- 建立门店销售预测模型（Prophet + 特征工程），MAPE 从人工经验的 23% 降到 9.4%，缺货率下降 5.8 个百分点。
- 设计会员分层 RFM 模型并落地到 CRM，定向触达的转化率是全量群发的 3.2 倍。

## 项目经历

### 用户流失预警模型 *2023*
#### 主导 · Python + XGBoost *覆盖 1200 万用户*

- 特征工程覆盖行为、内容、社交三类共 84 个特征，AUC 0.83，Top 10% 高危用户召回了实际流失量的 41%。
- 模型输出接入运营触达系统，被预警用户的挽回率相比对照组提升 12.7%。

## 教育背景

### 中国人民大学 *2018.09 - 2021.06*
#### 应用统计 · 硕士 *GPA 3.8/4.0*

### 武汉大学 *2014.09 - 2018.06*
#### 信息管理与信息系统 · 本科

## 其他

- 英语 IELTS 7.0，可独立阅读英文论文与撰写英文报告
- 公开分享：《别再用平均数骗自己》，公司内部数据周会，参与 200+ 人
