document.addEventListener("backoffice.index.shown", (e) => {
  setModalSize("modal-sm");
});

document.addEventListener("backoffice.cashes.shown", (e) => {
  clearModalSize();
});

document.addEventListener("backoffice.cash.updated", async (e) => {
  if (Object.hasOwn(e.detail, "backofficeCashesPath"))
    await htmx.ajax("GET", e.detail.backofficeCashesPath, { target: "#modal-body" });
});

document.addEventListener("backoffice.cash.deleted", async (e) => {
  if (Object.hasOwn(e.detail, "backofficeCashesPath"))
    await htmx.ajax("GET", e.detail.backofficeCashesPath, { target: "#modal-body" });
});

document.addEventListener("backoffice.cash.created", async (e) => {
  if (Object.hasOwn(e.detail, "backofficeCashesPath"))
    await htmx.ajax("GET", e.detail.backofficeCashesPath, { target: "#modal-body" });
});

document.addEventListener("backoffice.categories.shown", (e) => {
  setModalSize("modal-lg");
});

document.addEventListener("backoffice.category.updated", async (e) => {
  if (Object.hasOwn(e.detail, "backofficeCategoriesPath"))
    await htmx.ajax("GET", e.detail.backofficeCategoriesPath, { target: "#modal-body" });
});
