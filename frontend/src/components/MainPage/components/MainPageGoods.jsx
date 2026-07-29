
import {useNavigate} from "react-router";
import { memo } from "react";
import "./style.css"
import {config} from "../../../config"



const MainPageGoods = ({itemsList}) => {
    const {items, filteredItems} = itemsList
    const navigate = useNavigate();

    return (
        <div>
        {items && filteredItems  ? (
        <div className = "goods__wrapper">
            { items && !filteredItems ? (
                items?.map((item) => (
                <div className = "goods__item"
                key = {item.id}>
                    <img alt = "homa" src = {`${config.imgServerUrl}${item.item_picture}`}></img>
                    <div className ="goods-info__wrapper">
                        <div key = {item.id} className="good-main-info">
                            <p>{item.name}</p>
                            <p>{item.item_type}</p>
                            <p>{item.item_spec_id}</p>
                            {item.visible ? (
                                <p>Предмет не виден для других пользователей</p>
                            ) : (
                                <></>
                            )}
                            <button className ="goods-info__btn" onClick = {() => navigate(`/item/${item.id}`)}>купить</button>
                        </div>
                    </div>
                </div>
                    ))
                ) : (
                filteredItems?.map((item) => (
                    <div className = "goods__item"
                         key = {item.id}>
                        <img alt = "homa" src = {`${config.imgServerUrl}${item.item_picture}`}></img>
                        <div className ="goods-info__wrapper">
                            <div key = {item.id} className="goods-info">
                                <span className = "goods-info-name">{item.name}</span>
                                <span className = "goods-info-type">{item.item_type}</span>
                                <span className = "goods-info-article">{item.item_secondary_type}</span>
                                {item.visible ? (
                                    <></>
                                ) : (
                                    <span style = {{color:"red", fontWeight:600}}>Предмет виден только администратору</span>
                                )}
                                <button className ="goods-info__btn" onClick = {() => navigate(`/item/${item.id}`)}>УЗНАТЬ БОЛЬШЕ</button>
                            </div>
                        </div>
                    </div>
                    )))}
        </div>
            ):(
                <div>Товары отсутствуют</div>
        )}
        </div>
    )

}
export default memo(MainPageGoods);