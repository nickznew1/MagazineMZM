import {AnimatePresence, motion} from "framer-motion";
import {memo, useState} from "react";
import "./styles.css"
import {config} from "../../../config";





const CheckoutItems = ({checkout}) => {
    const inputs = checkout;
    const [activeSpec, setActiveSpec] = useState(false)
    const handleSpec =(id) =>{
        setActiveSpec(activeSpec === id ? null : id)
    }
    return (
        <div className = "checkout-items">
            {inputs.cartData.map((item, index) => (
                <div id = {index}  key={index} className = "checkout-cart-items">
                    <img alt = {item.name} src = {`${config.imgServerUrl}${item.item_picture}`}></img>
                    <span id ="name">{item.name}</span>
                    <span id ="item-type">Тип: {item.item_type}</span>
                    <span id = "item-secondary_type">Тип измерения: {item.item_secondary_type}</span>
                    <span id = "item-count"> Количество: {item.count}</span>
                    <button className = {activeSpec === item.item_spec_id ? "active" : ""} onClick = {()=>handleSpec(item.item_spec_id)} key ={index}>Показать выбранную спецификацию
                        <svg className = "filter-dropdown-icon" viewBox="0 0 24 24" width="16" height="16" stroke="currentColor"
                             strokeWidth="2" fill="none" strokeLinecap="round" strokeLinejoin="round">
                            <polyline points="6 9 12 15 18 9"></polyline>
                        </svg>
                    </button>
                    <AnimatePresence>
                        {activeSpec === item.item_spec_id &&
                            <motion.div
                                initial = {{ height:0, opacity:0}}
                                animate ={{height:"auto", opacity:1}}
                                exit={{height:0, opacity:0}}
                                transition={{duration: 0.3}}
                                className = "checkout-cart_item__specs-wrapper">
                                <div className = "checkout-cart_item__specs">
                                    {Object.entries(item.props).map(([key,value]) => (
                                        <>
                                            <p key ={key} id ="key">{key}:</p>
                                            <p key ={value} id ="value">{value}</p>
                                        </>
                                    ))}
                                </div>
                            </motion.div>
                        }
                    </AnimatePresence>
                </div>
            ))}
        </div>
    )
}

export default memo(CheckoutItems);