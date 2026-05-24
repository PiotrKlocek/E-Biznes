import { useState } from "react";

function Login() {

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  async function handleLogin(e) {
    e.preventDefault();

    const response = await fetch(
      "http://localhost:5000/api/login",
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({
          email,
          password
        })
      }
    );

    const data = await response.json();

    if (data.token) {
      localStorage.setItem("token", data.token);

      alert("Zalogowano");
    } else {
      alert(data.message);
    }
  }

  return (
    <div>
      <h1>Logowanie</h1>

      <form onSubmit={handleLogin}>

        <input
          type="email"
          placeholder="Email"
          onChange={(e) => setEmail(e.target.value)}
        />

        <br /><br />

        <input
          type="password"
          placeholder="Hasło"
          onChange={(e) => setPassword(e.target.value)}
        />

        <br /><br />

        <button type="submit">
          Zaloguj
        </button>
      </form>

      <br />

      <button
        onClick={() => {
          window.location.href =
            "http://localhost:5000/auth/google";
        }}
      >
        Login Google
      </button>

      <br /><br />

      <button
        onClick={() => {
          window.location.href =
            "http://localhost:5000/auth/github";
        }}
      >
        Login GitHub
      </button>

    </div>
  );
}

export default Login;