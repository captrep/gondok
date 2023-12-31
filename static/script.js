const base = window.location.href;

const addUrl = async (e) => {
  const form = document.getElementById("form");
  e.preventDefault();
  const formData = new FormData(form);
  let data = {};
  for (let key of formData.keys()) {
    data[key] = formData.get(key);
  }
  const status = document.getElementById("status");
  status.innerHTML = "Loading...";
  console.log(base)
  const response = await fetch(`${base}api/create`, {
    method: "POST",
    body: JSON.stringify(data),
  });
  const res = await response.json()
  if (response.status === 201) {
    status.innerHTML = `url created! you can open using link <a target="_blank" href="${base}${data.short_url}">${base}${data.short_url} </a>`;
  } else if (response.status === 409) {
    status.innerHTML = "Short url already exist. Please use other url";
  } else if (response.status === 400) {
    status.innerHTML = res.errors;
  } else {
    status.innerHTML = "something is wrong, try again later";
  }
};