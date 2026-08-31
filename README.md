<div align="center">

# M Mahesa

<img src="https://readme-typing-svg.demolab.com?font=Fira+Code&weight=600&size=20&duration=2800&pause=1000&color=00C2A8&center=true&vCenter=true&width=620&lines=Backend+%26+Network+Engineer;Go+%C2%B7+Python+%C2%B7+TypeScript+%C2%B7+PostgreSQL;Monitoring%2C+OLT%2FPON+automation%2C+and+NOC+tooling" alt="Backend & Network Engineer" />

<p>
  <a href="https://github.com/MMahesa"><img src="https://img.shields.io/badge/GitHub-MMahesa-181717?style=for-the-badge&logo=github&logoColor=white" alt="GitHub" /></a>
  <img src="https://img.shields.io/badge/Focus-Backend%20%26%20Infrastructure-0f172a?style=for-the-badge" alt="Focus" />
  <img src="https://img.shields.io/badge/Domain-ISP%20%2F%20NOC-14b8a6?style=for-the-badge" alt="Domain" />
</p>

I build the systems that keep a network operations team informed: monitoring back ends, device automation for OLT/PON fleets, and internal tools that turn raw telemetry into decisions.

</div>

---

## What I do

- **Backend services** — Go and Python services with PostgreSQL, designed around explicit state machines, idempotent workers, and verifiable rollbacks.
- **Network monitoring & automation** — SNMP/CLI collectors for routers and GPON/EPON OLTs (ZTE, HSGQ, HIOSO), alarm correlation, traffic history, and typed, preview-then-confirm device operations.
- **NOC tooling** — dashboards and ledgers that a shift operator can read at a glance: table-first UIs, honest freshness indicators, and no noisy warnings.
- **Operations discipline** — reproducible release packages, fail-fast deploy gates with automatic rollback, dated change notes, and incident write-ups with prevention items.

## Selected work

| Project | Stack | What it does |
| --- | --- | --- |
| **Network Monitor** *(private)* | Go · Vue 3 · PostgreSQL · SNMP/Telnet | Multi-vendor OLT/router monitoring for an ISP NOC: ONU lifecycle (power-down / laser-out / initial), PON status from ONU evidence, ZTE C620 typed operations with rollback, server-owned telemetry refresh, release harness with reproducible builds. |
| **Fiber Core Mapping** *(private)* | Django · GeoDjango · SvelteKit · MapLibre · Android | Fiber network mapping and field reporting: POP/ODC/ODP inventory, geospatial API, mobile client for technicians. |
| **Network Toolkit** *(private)* | Kotlin · Jetpack Compose | Android toolbox for field engineers: subnet & port utilities, encrypted credential vault, quick diagnostics. |
| [**port-scanner**](https://github.com/MMahesa/port-scanner) | Go | Concurrent TCP port scanner with bounded workers and clean output. |
| [**tools-registrasi**](https://github.com/MMahesa/tools-registrasi) | JavaScript | Registration utility with practical backend validation flow. |
| [**discord-lyrics**](https://github.com/MMahesa/discord-lyrics) | JavaScript | Discord bot that fetches and formats song lyrics through a public API. |

## Tech stack

<div align="center">
  <img src="https://skillicons.dev/icons?i=go,python,ts,js,vue,svelte,django,kotlin,postgres,redis,docker,linux,nginx,git,github,vscode&perline=8" alt="Tech stack" />
</div>

<details>
<summary><b>More detail</b></summary>
<br />

| Area | Tools |
| --- | --- |
| Languages | Go, Python, TypeScript/JavaScript, Kotlin, SQL |
| Back end | net/http, pgx, Django/DRF, Node.js |
| Front end | Vue 3 + TanStack Query, SvelteKit, ECharts, MapLibre |
| Data | PostgreSQL 16, Redis |
| Network | SNMP (v2c/v3), syslog/trap receivers, RouterOS & VyOS, GPON/EPON OLT CLI (ZTE C620, HSGQ, HIOSO) |
| Ops | systemd, Docker, nginx, Tailscale, Playwright/vitest, Go test, reproducible release packaging |

</details>

## How I work

```text
Validate → Preview → Confirm → Execute → Verify
```

- Every production change ships with a backup, a hash-pinned baseline, a health gate, and an automatic rollback path.
- Device writes are typed and bounded — no arbitrary CLI, no guessed commands, evidence before "online".
- Findings and incidents are written down the same day, with root cause and a prevention item.

## GitHub activity

<div align="center">
  <img src="https://raw.githubusercontent.com/MMahesa/MMahesa/main/github-metrics.svg" alt="GitHub metrics" />
  <br /><br />
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/MMahesa/MMahesa/output/pacman-contribution-graph-dark.svg">
    <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/MMahesa/MMahesa/output/pacman-contribution-graph.svg">
    <img alt="Contribution graph" src="https://raw.githubusercontent.com/MMahesa/MMahesa/output/pacman-contribution-graph.svg" />
  </picture>
</div>

## Contact

<div align="center">
  <a href="https://github.com/MMahesa"><img src="https://img.shields.io/badge/GitHub-@MMahesa-181717?style=flat-square&logo=github" alt="GitHub" /></a>
  <a href="https://github.com/MMahesa?tab=repositories"><img src="https://img.shields.io/badge/Repositories-browse-0f172a?style=flat-square&logo=git&logoColor=white" alt="Repositories" /></a>
</div>

<br />

<div align="center">
  <sub>"Hal-hal besar dimulai dari hal kecil yang ditekuni."</sub>
</div>
