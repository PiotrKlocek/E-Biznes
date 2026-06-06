import { useState } from "react";
import api from "../api/axios";
import { useShop } from "../context/ShopContext";
import { useToast } from "../context/ToastContext";

function Payments() {
  const { total, clearCart } = useShop();
  const { addToast } = useToast();
  const [form, setForm] = useState({ fullName: "", email: "" });

  const handleChange = (e) => setForm((p) => ({ ...p, [e.target.name]: e.target.value }));

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await api.post("/payments", { ...form, amount: total });
      addToast("Płatność przyjęta!");
      clearCart();
      setForm({ fullName: "", email: "" });
    } catch {
      addToast("Błąd płatności", "error");
    }
  };

  return (
    <div className="container">
      <h1>Płatność</h1>
      <div className="form-box">
        <form onSubmit={handleSubmit}>
          <div className="form-group">
            <label>Imię i nazwisko</label>
            <input type="text" name="fullName" value={form.fullName} onChange={handleChange} required />
          </div>
          <div className="form-group">
            <label>Email</label>
            <input type="email" name="email" value={form.email} onChange={handleChange} required />
          </div>
          <div className="form-group">
            <label>Do zapłaty</label>
            <p><strong>{total.toFixed(2)} zł</strong></p>
          </div>
          <button className="btn-success" type="submit">Zapłać</button>
        </form>
      </div>
    </div>
  );
}

export default Payments;
