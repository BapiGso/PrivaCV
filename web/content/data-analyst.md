<style>
/* ===========================================================
   Banking — 复刻 CTAN 上 moderncv 宏包的 banking 样式
   （\moderncvstyle{banking}\moderncvcolor{blue}）

   签名细节：姓名齐左、细体大字，头部下面一条粗强调色横线，
   章节标题着色全大写、下面一条发丝线。克制但一眼认得出。
   =========================================================== */

:root {
    --font-body: var(--font-sans);

    --accent: #2b5c8a;          /* moderncv blue，可换 #7a2e2e / #2f6b46 */
    --fs: 10.1pt;
    --lh: 1.48;
    --fs-name: 24pt;
    --fs-sec: 10.8pt;

    --page-x: 17mm;
    --page-y: 15mm;
    --sp-sec: 5.4mm;
    --sp-item: 3.2mm;

    --ink: #232323;
    --muted: #616161;
    --rule-soft: #d5dde5;
    --tint: rgba(43, 92, 138, 0.09);
}

h1 {
    font-weight: 300;
    letter-spacing: -0.008em;
    margin: 0 0 1.6mm;
}

h1 strong {
    font-weight: 600;
    color: var(--accent);
}

h1 + p {
    font-size: calc(var(--fs) * 0.9);
    line-height: 1.55;
    margin: 0 0 var(--sp-sec);
    padding-bottom: 3.2mm;
    border-bottom: 2.4pt solid var(--accent);
}

h2 {
    color: var(--accent);
    font-weight: 700;
    text-transform: uppercase;
    letter-spacing: 0.09em;
    padding-bottom: 1mm;
    border-bottom: 0.5pt solid var(--rule-soft);
    margin: var(--sp-sec) 0 2mm;
}

h3 { font-weight: 600; }

h4 { color: var(--muted); }

h3 > .x {
    color: var(--accent);
    font-weight: 500;
}

li::marker { color: var(--accent); }

a { color: var(--accent); }
h1 + p a { color: inherit; }

code { color: var(--accent); }

blockquote {
    border-left-color: var(--accent);
    font-style: italic;
}

td:first-child { color: var(--accent); }
</style>

# 陈亦然

数据分析师 · 4 年经验 ⋄ 北京 ⋄ 137-0000-0000 ⋄ chenyiran@example.com

## 技能

| 类别 | 内容 |
|:--|:--|
| 分析 | SQL（复杂窗口函数、性能调优）、Python（pandas / statsmodels / scikit-learn） |
| 实验 | A/B 实验设计、假设检验、因果推断（DID、PSM）、样本量估算 |
| 可视化 | Tableau、Superset、Matplotlib、ECharts |
| 数据栈 | Hive、Spark、ClickHouse、dbt、Airflow |

## 工作经历

### 某内容社区 · 增长数据组 <span>2022.08 - 至今</span>
#### 数据分析师 <span>北京</span>

- 搭建新用户 7 日留存的归因分析框架，定位到「首次关注引导」是留存的关键路径；据此推动的产品改版使 D7 留存从 24.1% 提升到 29.6%。
- 重构核心指标口径与看板，把散落在 4 个团队的 30 余个「日活」定义收敛为 1 套，月度对账争议从平均 11 次降到 0。
- 负责社区激励预算的效果评估，用 DID 方法识别出 2 个 ROI 为负的补贴场景，全年节省预算约 480 万元。
- 建立 A/B 实验平台的元分析机制，复盘全年 213 场实验，发现 37% 的实验因样本量不足无法得出结论，推动实验准入规范落地。

### 某零售集团 <span>2021.06 - 2022.07</span>
#### 商业分析师 <span>北京</span>

- 建立门店销售预测模型（Prophet + 特征工程），MAPE 从人工经验的 23% 降到 9.4%，缺货率下降 5.8 个百分点。
- 设计会员分层 RFM 模型并落地到 CRM，定向触达的转化率是全量群发的 3.2 倍。

## 项目经历

### 用户流失预警模型 <span>2023</span>
#### 主导 · Python + XGBoost <span>覆盖 1200 万用户</span>

- 特征工程覆盖行为、内容、社交三类共 84 个特征，AUC 0.83，Top 10% 高危用户召回了实际流失量的 41%。
- 模型输出接入运营触达系统，被预警用户的挽回率相比对照组提升 12.7%。

## 教育背景

### 中国人民大学 <span>2018.09 - 2021.06</span>
#### 应用统计 · 硕士 <span>GPA 3.8/4.0</span>

### 武汉大学 <span>2014.09 - 2018.06</span>
#### 信息管理与信息系统 · 本科 <span></span>

## 其他

- 英语 IELTS 7.0，可独立阅读英文论文与撰写英文报告
- 公开分享：《别再用平均数骗自己》，公司内部数据周会，参与 200+ 人
