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

:root {
    --font-body: var(--font-sans);

    --accent: #3d5a52;
    --fs: 9.6pt;
    --lh: 1.5;
    --fs-name: 24pt;
    --fs-sec: 9.8pt;

    --page-x: 16mm;
    --page-y: 15mm;
    --sp-sec: 5.2mm;
    --sp-item: 3.2mm;
    --sp-line: 1.2mm;

    --ink: #232725;
    --muted: #6b716e;
    --rule-soft: #dfe3e1;
    --tint: rgba(61, 90, 82, 0.09);
}

/* --- 通栏头部 --- */
h1 {
    font-weight: 600;
    letter-spacing: -0.012em;
    margin: 0 0 1.4mm;
}

h1 + p {
    font-size: calc(var(--fs) * 0.94);
    line-height: 1.55;
    color: var(--accent);
    letter-spacing: 0.03em;
    margin: 0 0 5.4mm;
    padding-bottom: 3.4mm;
    border-bottom: 1.6pt solid var(--accent);
}

/* --- 侧栏 34%，中间一条贯穿到底的发丝线 --- */
.cv-grid {
    display: grid;
    grid-template-columns: 34% 1fr;
    column-gap: 9mm;
    align-items: start;
    background: linear-gradient(var(--rule-soft), var(--rule-soft)) no-repeat;
    background-size: 0.4pt 100%;
    background-position: calc(34% + 4.5mm) 0;
}

.cv-grid > .main { padding-left: 1mm; }

h2 {
    color: var(--accent);
    font-weight: 700;
    text-transform: uppercase;
    letter-spacing: 0.13em;
    font-size: var(--fs-sec);
    margin: var(--sp-sec) 0 1.8mm;
}

.main h2 {
    padding-bottom: 0.9mm;
    border-bottom: 0.5pt solid var(--rule-soft);
}

h3 { font-weight: 600; }

h4 {
    color: var(--muted);
    font-style: italic;
}

[data-split] { gap: 3mm; }
[data-split] > .x { font-size: 0.9em; }

li::marker { color: var(--accent); }

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
    color: var(--muted);
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

a { color: var(--accent); }
h1 + p a { color: inherit; }

code { color: var(--accent); border-radius: 1mm; }

blockquote { border-left-color: var(--accent); }

td:first-child { width: 20mm; color: var(--accent); }
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

### Regional Fintech Platform <span>2021 – present</span>
#### Senior Software Engineer, Payments Core <span>Singapore</span>

- Led the migration of the settlement ledger from a single Postgres instance to a sharded, event-sourced design; sustained 12k TPS at peak with zero reconciliation discrepancies across 18 months.
- Introduced idempotency keys and an outbox pattern across 11 services, eliminating a class of duplicate-charge incidents that previously cost roughly SGD 40k per quarter in refunds.
- Cut p99 authorization latency from 840 ms to 180 ms by replacing synchronous risk checks with a pre-computed decision cache.
- Mentored 3 engineers; two were promoted within 18 months.

### Regional E-Commerce Company <span>2018 – 2021</span>
#### Software Engineer, Messaging Infrastructure <span>Singapore</span>

- Built a push notification pipeline handling 400M messages/day with at-least-once delivery and per-tenant rate limiting.
- Reduced infrastructure cost by 38% by replacing a fan-out-on-write design with fan-out-on-read for low-engagement segments.
- Ran the on-call rotation redesign: alert volume dropped 64% after pruning non-actionable alerts and adding SLO-based paging.

## Selected Projects

### Open-source rate limiter <span>2022 – present</span>
#### Author · Go <span>[github.com/example/limiter](https://github.com/example/limiter)</span>

- Distributed token bucket backed by Redis, sub-millisecond overhead, 2.1k stars, used in production by several fintech teams.

## Talks

- *Idempotency Is a Product Decision* — GopherCon SG, 2024
- *What Sharding a Ledger Actually Costs You* — QCon Asia, 2023

</div>
</div>
