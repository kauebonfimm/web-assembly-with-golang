const go = new Go();

WebAssembly.instantiateStreaming(
  fetch("assets/assembly.wasm"),
  go.importObject
).then((result) => {
  go.run(result.instance);
  console.log("Wasm loaded successfully");
});

function generatePassword() {
  const passwordLength = document.getElementById("length").value;

  if (passwordLength < 1) {
    alert("Password length must be greater than 0");
    return;
  }

  if (passwordLength > 1000) {
    alert("Password length must be less than 1000");
    return;
  }

  const hasLetters = document.getElementById("letters").checked;
  const hasNumbers = document.getElementById("numbers").checked;
  const hasSymbols = document.getElementById("symbols").checked;
  const removeCaracters = document.getElementById("ignoreCaracters").value;

  const password = generate(
    parseInt(passwordLength),
    hasLetters,
    hasNumbers,
    hasSymbols,
    removeCaracters
  );

  document.getElementById("password").innerText = password;

  console.log("Password generated successfully");
}
