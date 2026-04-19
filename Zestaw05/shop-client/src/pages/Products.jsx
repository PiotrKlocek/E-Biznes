import { useEffect, useState } from "react";
import api from "../api/axios";
import { useShop } from "../context/ShopContext";

function Products() {
  const [products, setProducts] = useState([]);
  const { addToCart } = useShop();

  useEffect(() => {
    api
      .get("/products")
      .then((res) => setProducts(res.data))
      .catch((err) => console.error("Błąd pobierania produktów:", err));
  }, []);

  return (
    <div className="container">
      <h1 className="page-title">Produkty</h1>

      <div className="products-grid">
        {products.map((product) => (
          <div className="card" key={product.id}>
            <h3>{product.name}</h3>
            <p className="price">{product.price.toFixed(2)} zł</p>
            <button className="btn-primary" onClick={() => addToCart(product)}>
              Dodaj do koszyka
            </button>
          </div>
        ))}
      </div>
    </div>
  );
}

export default Products;