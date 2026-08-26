<style>
/* ===========================================================
   Sidebar — 复刻 AltaCV（github.com/liantze/AltaCV）
   就是那份被叫作「Marissa Mayer 简历」的排版。

   原版侧栏带大片底色，这里换成一条发丝线：留白分栏比色块
   耐看，打印出来也不会糊成一团。侧栏放联系方式、技能、语言，
   主栏放经历。

   需要正文里有 <div class="cv-grid"> + .sidebar / .main，
   没有时会自动退化成单栏。
   =========================================================== */

body {
    padding: 15mm 16mm;
    font-family: Inter, "Source Sans 3", "Helvetica Neue", Helvetica, Arial,
                 "Segoe UI", Roboto, "PingFang SC", "Source Han Sans SC",
                 "Microsoft YaHei", sans-serif;
    font-size: 9.6pt;
    line-height: 1.5;
    color: #232725;
}

/* --- 通栏头部 --- */
h1 {
    font-weight: 600;
    font-size: 24pt;
    letter-spacing: -0.012em;
    line-height: 1.15;
    margin: 0 0 1.4mm;
}

h1 + p {
    font-size: 9.02pt;
    line-height: 1.55;
    color: #3d5a52;
    letter-spacing: 0.03em;
    margin: 0 0 5.4mm;
    padding-bottom: 3.4mm;
    border-bottom: 1.6pt solid #3d5a52;
}

/* --- 侧栏 34%，中间一条贯穿到底的发丝线 --- */
.cv-grid {
    display: grid;
    grid-template-columns: 34% 1fr;
    column-gap: 9mm;
    align-items: start;
    background: linear-gradient(#dfe3e1, #dfe3e1) no-repeat;
    background-size: 0.4pt 100%;
    background-position: calc(34% + 4.5mm) 0;
}

.cv-grid > .main { padding-left: 1mm; }

h2 {
    color: #3d5a52;
    font-weight: 700;
    text-transform: uppercase;
    letter-spacing: 0.13em;
    font-size: 9.8pt;
    line-height: 1.25;
    margin: 5.2mm 0 1.8mm;
}

/* 两栏各自的头一个章节标题贴着栏顶 */
.sidebar h2:first-of-type,
.main h2:first-of-type { margin-top: 0; }

.main h2 {
    padding-bottom: 0.9mm;
    border-bottom: 0.5pt solid #dfe3e1;
}

h3 {
    font-weight: 600;
    font-size: 9.6pt;
    line-height: 1.3;
    margin: 3.2mm 0 0;
}

h4 {
    color: #6b716e;
    font-style: italic;
    font-weight: 400;
    font-size: 8.83pt;
    line-height: 1.35;
    margin: 0.4mm 0 1.2mm;
}

/* 时间和地点：斜体写在标题末尾，浮到右边 */
h3 em, h4 em {
    float: right;
    font-weight: 400;
    font-size: 0.9em;
    font-variant-numeric: tabular-nums;
}
h3 em { font-style: normal; }

/* 左栏只有 34% 宽，浮动会把标题挤散，在那儿就当普通斜体 */
.sidebar h3 em, .sidebar h4 em { float: none; }

p { margin: 0 0 1.2mm; }
ul, ol { margin: 1.2mm 0 1.8mm; padding-left: 4.6mm; }
li { margin: 0 0 0.9mm; }
li::marker { color: #3d5a52; font-size: 0.9em; }

/* 侧栏是一张清单，不是正文：无项目符号、行距更紧 */
.sidebar { font-size: 0.97em; }
.sidebar ul { list-style: none; padding-left: 0; }
.sidebar li { margin-bottom: 1mm; }
.sidebar p { margin-bottom: 1mm; }

.sidebar blockquote {
    border: 0;
    padding: 0;
    margin: 0.4mm 0 1.4mm;
    font-size: 0.94em;
    color: #6b716e;
}

/* 侧栏顶部放头像时自动裁成圆形 */
.sidebar img {
    display: block;
    width: 30mm;
    height: 30mm;
    object-fit: cover;
    border-radius: 50%;
    margin: 0 auto 4mm;
}

a { color: #3d5a52; text-decoration: none; }
h1 + p a { color: inherit; }

/* 行内 code 当技能标签用：浅底圆角，字体跟正文走 */
code {
    font-family: inherit;
    font-size: 0.92em;
    padding: 0.3mm 1.4mm;
    border-radius: 1mm;
    background: rgba(61, 90, 82, 0.09);
    white-space: nowrap;
    color: #3d5a52;
}

blockquote {
    margin: 1.2mm 0;
    padding: 0 0 0 3.5mm;
    border-left: 1.2pt solid #3d5a52;
    color: #6b716e;
}
</style>

# Wei Chen

Senior Software Engineer · Distributed Systems ⋄ Singapore ⋄ +65 8000 0000 ⋄ wei.chen@example.com

<div class="cv-grid">
<div class="sidebar">

## Contact

- wei.chen@example.com
- +65 8000 0000
- Singapore (PR, no visa needed)
- [github.com/example](https://github.com/example)
- [linkedin.com/in/example](https://linkedin.com/in/example)

## Skills

- Go, Java, Python, TypeScript
- PostgreSQL, Redis, Kafka
- Kubernetes, Terraform, AWS
- gRPC, GraphQL, REST
- Prometheus, Grafana, OpenTelemetry

## Education

**National University of Singapore**

> B.Comp. in Computer Science
> 2014 – 2018, First Class Honours

## Languages

- English — fluent (working language)
- Mandarin — native
- Malay — conversational

## Certifications

- AWS Solutions Architect – Professional
- CKA (Certified Kubernetes Administrator)

</div>
<div class="main">

## Summary

Backend engineer with 7 years building payment and messaging infrastructure at scale. I like problems where correctness and throughput are both non-negotiable, and I care more about systems that fail predictably than systems that are merely fast.

## Experience

### Regional Fintech Platform *2021 – present*
#### Senior Software Engineer, Payments Core *Singapore*

- Led the migration of the settlement ledger from a single Postgres instance to a sharded, event-sourced design; sustained 12k TPS at peak with zero reconciliation discrepancies across 18 months.
- Introduced idempotency keys and an outbox pattern across 11 services, eliminating a class of duplicate-charge incidents that previously cost roughly SGD 40k per quarter in refunds.
- Cut p99 authorization latency from 840 ms to 180 ms by replacing synchronous risk checks with a pre-computed decision cache.
- Mentored 3 engineers; two were promoted within 18 months.

### Regional E-Commerce Company *2018 – 2021*
#### Software Engineer, Messaging Infrastructure *Singapore*

- Built a push notification pipeline handling 400M messages/day with at-least-once delivery and per-tenant rate limiting.
- Reduced infrastructure cost by 38% by replacing a fan-out-on-write design with fan-out-on-read for low-engagement segments.
- Ran the on-call rotation redesign: alert volume dropped 64% after pruning non-actionable alerts and adding SLO-based paging.

## Selected Projects

### Open-source rate limiter *2022 – present*
#### Author · Go *[github.com/example/limiter](https://github.com/example/limiter)*

- Distributed token bucket backed by Redis, sub-millisecond overhead, 2.1k stars, used in production by several fintech teams.

## Talks

- *Idempotency Is a Product Decision* — GopherCon SG, 2024
- *What Sharding a Ledger Actually Costs You* — QCon Asia, 2023

</div>
</div>
