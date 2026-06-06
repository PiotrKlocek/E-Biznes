import { useEffect, useState } from "react";
import api from "../api/axios";
import { useShop } from "../context/ShopContext";
import { useToast } from "../context/ToastContext";

function Products() {
  const [products, setProducts] = useState([]);
  const { addToCart } = useShop();
  const { addToast } = useToast();

  useEffect(() => {
    api.get("/products")
      .then((res) => setProducts(res.data))
      .catch(() => addToast("Błąd pobierania produktów", "error"));
  }, []);

  return (
    <div className="container">
      <h1>Produkty</h1>
      <div className="products-grid">
        {products.map((p) => (
          <div className="card" key={p.id}>
            <h3>{p.name}</h3>
            <p className="price">{p.price.toFixed(2)} zł</p>
            <button className="btn-primary" onClick={() => { addToCart(p); addToast(`${p.name} dodany`); }}>
              Dodaj do koszyka
            </button>
          </div>
        ))}
      </div>
    </div>
  );
}

export default Products;
