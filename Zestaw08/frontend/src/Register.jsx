import { useState } from "react";

function Register() {

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  async function handleRegister(e) {
    e.preventDefault();

    const response = await fetch(
      "http://localhost:5000/api/register",
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

    alert(data.message);
  }

  return (
    <div>

      <h1>Rejestracja</h1>

      <form onSubmit={handleRegister}>

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
          Zarejestruj
        </button>

      </form>

    </div>
  );
}

export default Register;