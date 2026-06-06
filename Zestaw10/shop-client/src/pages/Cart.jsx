import api from "../api/axios";
import { useShop } from "../context/ShopContext";
import { useToast } from "../context/ToastContext";

function Cart() {
  const { cart, removeFromCart, total } = useShop();
  const { addToast } = useToast();

  const sendCart = async () => {
    try {
      await api.post("/cart", { items: cart });
      addToast("Koszyk zapisany");
    } catch {
      addToast("Błąd zapisywania", "error");
    }
  };

  if (cart.length === 0) {
    return (
      <div className="container">
        <h1>Koszyk</h1>
        <p className="empty">Koszyk jest pusty</p>
      </div>
    );
  }

  return (
    <div className="container">
      <h1>Koszyk</h1>
      {cart.map((item, i) => (
        <div className="cart-item" key={i}>
          <div>
            <strong>{item.name}</strong> – {item.price.toFixed(2)} zł
          </div>
          <button className="btn-danger" onClick={() => removeFromCart(i)}>Usuń</button>
        </div>
      ))}
      <div className="total-box">
        <h2>Suma: {total.toFixed(2)} zł</h2>
        <button className="btn-success" onClick={sendCart}>Zapisz koszyk</button>
      </div>
    </div>
  );
}

export default Cart;
