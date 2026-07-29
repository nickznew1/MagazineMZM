import "./styles.css"
import {useNavigate} from "react-router";
import "./styles.css"
import {memo} from "react";
import {config} from "../../../config";

const ProductInfo = ({itemCard}) =>{
    const item = itemCard;
    const navigate = useNavigate();
    return (
        <>
        <div className="product-info__img-wrapper">
            <div className="product-info__img">
                <img alt = "homa" src ={`${config.imgServerUrl}${item.itemData.item_picture}`}></img>
            </div>
        </div>
    <div className ="product-info-description-wrapper">
        <div className="product-info-description">
            <h3>Краткие характеристики:</h3>
            <p>{item.itemData.item_short_description}</p>
        </div>
    </div>
    <div className ="product-price-info-wrapper">
        <div className ="product-price-info">
            <p className = "name">{item.itemData.name}</p>
            <p className = "price">Цена: от {item.itemData.price} Р</p>
            <button className="goods-info__btn" onClick={() => navigate(`/constructor/${item.itemData.id}`)}>Перейти в конструктор</button>
        </div>
    </div>
        </>
    )
}

export default memo(ProductInfo);