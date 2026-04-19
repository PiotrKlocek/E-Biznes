import api from "../api/axios";
import { useShop } from "../context/ShopContext";

function Cart() {
  const { cart, removeFromCart, total } = useShop();

  const sendCart = async () => {
    try {
      const res = await api.post("/cart", { items: cart });
      alert(res.data.message);
    } catch (err) {
      console.error("Błąd wysyłania koszyka:", err);
    }
  };

  return (
    <div className="container">
      <h1 className="page-title">Koszyk</h1>

      {cart.length === 0 ? (
        <div className="empty-box">
          <p>Koszyk jest pusty</p>
        </div>
      ) : (
        <>
          <div className="cart-list">
            {cart.map((item, index) => (
              <div className="cart-item" key={index}>
                <div className="cart-item-info">
                  <h3>{item.name}</h3>
                  <p>{item.price.toFixed(2)} zł</p>
                </div>

                <button
                  className="btn-danger"
                  onClick={() => removeFromCart(index)}
                >
                  Usuń
                </button>
              </div>
            ))}
          </div>

          <div className="total-box">
            <h2>Suma: {total.toFixed(2)} zł</h2>
            <button className="btn-success" onClick={sendCart}>
              Zapisz koszyk
            </button>
          </div>
        </>
      )}
    </div>
  );
}

export default Cart;