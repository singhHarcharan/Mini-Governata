(() => {
  const phases = [
    [0, "Orientation & tracing", "phase-00-orientation"],
    [1, "React workspace", "phase-01-frontend"],
    [2, "Go foundations", "phase-02-go-foundations"],
    [3, "Go HTTP API", "phase-03-go-api"],
    [4, "Interfaces & boundaries", "phase-04-interfaces"],
    [5, "PostgreSQL & GORM", "phase-05-postgres-gorm"],
    [6, "Auth, tenants & RBAC", "phase-06-auth-tenancy"],
    [7, "Workflow, audit & SLA", "phase-07-workflow-audit"],
    [8, "NATS & service split", "phase-08-nats-services"],
    [9, "Reliable events", "phase-09-reliable-events"],
    [10, "CEL policies", "phase-10-cel-policies"],
    [11, "Storage & workers", "phase-11-storage-workers"],
    [12, "Production proof", "phase-12-production"]
  ];
  const root = document.body.dataset.root || "..";
  const current = Number(document.body.dataset.phase ?? -1);
  const storageKey = "mini-governata-course-progress-v1";
  let state = {};
  try { state = JSON.parse(localStorage.getItem(storageKey) || "{}"); } catch { state = {}; }

  const phaseChecks = (phase) => Array.from(document.querySelectorAll(`[data-check^="p${phase}-"]`));
  const isDone = (phase) => {
    const checks = phase === current ? phaseChecks(phase) : [];
    if (checks.length) return checks.every(input => state[input.dataset.check]);
    return Boolean(state[`phase-${phase}-complete`]);
  };

  const nav = document.querySelector("[data-course-nav]");
  if (nav) {
    nav.className = "course-nav";
    nav.innerHTML = `
      <a class="brand" href="${root}/index.html">Mini Governata Academy</a>
      <p class="brand-subtitle">Learn the system by building it.</p>
      <div class="nav-progress"><div class="nav-progress-row"><span>Course progress</span><strong data-total-progress>0%</strong></div><div class="progress-track"><div class="progress-fill" data-progress-fill></div></div></div>
      <h2>Phase guides</h2>
      <nav class="course-nav-list" aria-label="Course phases">
        ${phases.map(([n, title, folder]) => `<a class="course-nav-link" ${n === current ? 'aria-current="page"' : ''} href="${root}/${folder}/index.html"><span class="nav-number">${n}</span><span>${title}</span><span class="nav-check" data-nav-check="${n}">${isDone(n) ? "✓" : ""}</span></a>`).join("")}
      </nav>
      <a class="nav-back" href="${root}/../mini-governata-learning-roadmap.html">← Master roadmap</a>`;
  }

  const updateProgress = () => {
    if (current >= 0) {
      const checks = phaseChecks(current);
      const complete = checks.length > 0 && checks.every(input => input.checked);
      state[`phase-${current}-complete`] = complete;
      const mark = document.querySelector(`[data-nav-check="${current}"]`);
      if (mark) mark.textContent = complete ? "✓" : "";
      const localLabel = document.querySelector("[data-phase-progress]");
      if (localLabel) {
        const done = checks.filter(input => input.checked).length;
        localLabel.textContent = checks.length ? `${done}/${checks.length} complete` : "Start when ready";
      }
    }
    const completed = phases.filter(([n]) => isDone(n)).length;
    const percent = Math.round(completed / phases.length * 100);
    document.querySelectorAll("[data-total-progress]").forEach(el => el.textContent = `${percent}%`);
    document.querySelectorAll("[data-progress-fill]").forEach(el => el.style.width = `${percent}%`);
    localStorage.setItem(storageKey, JSON.stringify(state));
  };

  document.querySelectorAll("[data-check]").forEach(input => {
    input.checked = Boolean(state[input.dataset.check]);
    input.addEventListener("change", () => {
      state[input.dataset.check] = input.checked;
      updateProgress();
    });
  });

  document.querySelectorAll("pre").forEach(pre => {
    const wrap = document.createElement("div");
    wrap.className = "code-wrap";
    pre.parentNode.insertBefore(wrap, pre);
    wrap.appendChild(pre);
    const button = document.createElement("button");
    button.className = "copy-code";
    button.type = "button";
    button.textContent = "Copy";
    button.addEventListener("click", async () => {
      await navigator.clipboard.writeText(pre.innerText);
      button.textContent = "Copied";
      setTimeout(() => button.textContent = "Copy", 1200);
    });
    wrap.appendChild(button);
  });

  updateProgress();
})();
