const API = "http://localhost:8080"

const status = document.getElementById("status")
const otpValue = document.getElementById("otpValue")
const topbar = document.getElementById("topbar")
const accountName = document.getElementById("accountName")

function setStatus(msg, type = "info") {
  status.className = `status ${type}`
  status.innerText = msg
}

function hideAll() {
  choiceBox.style.display = "none"
  registerBox.classList.add("hidden")
  loginBox.classList.add("hidden")
  otpBox.classList.add("hidden")
  successBox.classList.add("hidden")
}


function showRegister() {
  hideAll()
  registerBox.classList.remove("hidden")
  setStatus("Create a new account", "info")
}

function showLogin() {
  hideAll()
  loginBox.classList.remove("hidden")
  setStatus("Login with your credentials", "info")
}

function registerUser() {

  if (!checkPasswordStrength(rpass.value)) {
    setStatus("Please use a stronger password", "error")
    return
  }
  setStatus("Registering...", "info")
  fetch(API + "/register", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      username: ruser.value.trim(),
      password: rpass.value.trim()
    })
  })
  .then(r => {
    if (!r.ok) return r.text().then(t => { throw t })
    return r.text()
  })
  .then(msg => {
    setStatus(msg, "success")
    setTimeout(showLogin, 1200)
  })
  .catch(err => setStatus(err, "error"))
}

function checkPasswordStrength(pwd) {
  const hint = document.getElementById("passwordHint")

  const strong =
    pwd.length >= 8 &&
    /[A-Z]/.test(pwd) &&
    /[a-z]/.test(pwd) &&
    /[0-9]/.test(pwd)

  if (pwd.length === 0) {
    hint.className = "password-hint"
    return false
  }

  if (strong) {
    hint.className = "password-hint strong"
    hint.innerText = "Strong password ✔"
    return true
  } else {
    hint.className = "password-hint weak"
    hint.innerText = "Weak password ✖"
    return false
  }
}


function loginUser() {
  setStatus("Verifying credentials...", "info")

  fetch(API + "/login", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      username: luser.value.trim(),
      password: lpass.value.trim()
    })
  })
  .then(r => {
    if (!r.ok) return r.text().then(t => { throw t })
    return r.json()
  })
  .then(data => {
    otpValue.innerText = data.otp
    hideAll()
    otpBox.classList.remove("hidden")
    setStatus(data.message, "success")
  })
  .catch(err => setStatus(err, "error"))
}

function verifyOtp() {
  setStatus("Verifying OTP...", "info")

  fetch(API + "/verify-otp", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      username: luser.value.trim(),
      otp: otp.value.trim()
    })
  })
  .then(r => {
    if (!r.ok) return r.text().then(t => { throw t })
    return r.text()
  })
  .then(msg => {
    hideAll()
    successBox.classList.remove("hidden")
    setStatus(msg, "success")

    topbar.classList.remove("hidden")
    accountName.innerText = luser.value
  })
  .catch(err => setStatus(err, "error"))
}


function logout() {
  // Clear inputs
  ruser.value = ""
  rpass.value = ""
  luser.value = ""
  lpass.value = ""
  otp.value = ""
  otpValue.innerText = "----"

  // Hide everything
  hideAll()
  successBox.classList.add("hidden")
  topbar.classList.add("hidden")

  // Show start screen
  choiceBox.style.display = "block"
  setStatus("You have logged out safely", "info")
}

rpass.addEventListener("input", e => checkPasswordStrength(e.target.value))
