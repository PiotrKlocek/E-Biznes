import { useState } from "react";
import api from "../api/axios";
import { useShop } from "../context/ShopContext";

function Payments() {
  const { total, clearCart } = useShop();
  const [form, setForm] = useState({
    fullName: "",
    email: "",
  });

  const handleChange = (e) => {
    setForm((prev) => ({
      ...prev,
      [e.target.name]: e.target.value,
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const res = await api.post("/payments", {
        ...form,
        amount: total,
      });

      alert(res.data.message);
      clearCart();
      setForm({
        fullName: "",
        email: "",
      });
    } catch (err) {
      console.error("Błąd płatności:", err);
    }
  };

  return (
    <div className="container">
      <h1 className="page-title">Płatności</h1>

      <div className="form-box">
        <form onSubmit={handleSubmit}>
          <div className="form-group">
            <input
              type="text"
              name="fullName"
              placeholder="Imię i nazwisko"
              value={form.fullName}
              onChange={handleChange}
            />
          </div>

          <div className="form-group">
            <input
              type="email"
              name="email"
              placeholder="Email"
              value={form.email}
              onChange={handleChange}
            />
          </div>

          <div className="form-group">
            <p>Do zapłaty: <strong>{total.toFixed(2)} zł</strong></p>
          </div>

          <button className="btn-success" type="submit">
            Zapłać
          </button>
        </form>
      </div>
    </div>
  );
}

export default Payments;