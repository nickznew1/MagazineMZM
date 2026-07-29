import "./styles.css"
import {memo, useState} from "react";
import {config} from "../../../config";

const ProductInfoHeader = ({itemCard, userRole}) => {

    const visibleStatus = {
        "Show":true,
        "Not show":false,
    }

    const [visible, setVisible] = useState(itemCard.itemData.visible ? visibleStatus["Show"] : visibleStatus["Not show"] );
    const [visibleMessage, setVisibleMessage] = useState(!itemCard.itemData.visible)

    const handleChangeVisible = (status) =>{
        if (status){
            setVisible(true);
            setVisibleMessage(false)
        } else {
            setVisible(false)
            setVisibleMessage(true)
        }
        handleVisible(status);
    }



    const handleVisible = async (status) => {
        const payload = {
            visible : status,
        }

        try {
            const response = await fetch (`${config.apiUrl}/admin/visible/${itemCard.itemData.id}`,{
                method :"POST",
                headers:{
                    "Content-Type": "application/json",
                },
                body:JSON.stringify(payload)
            })


        } catch (err){
            setVisible(false)
            throw new Error(err)
        }


    }
    console.log(visible)



    return(
        <div className ="product-info-header-wrapper">
            <div className="product-info-name">
                <h1>{itemCard.itemData.name}</h1>
                {userRole === "admin" && visibleMessage && (
                    <h1 style = {{color: "red"}}>Товар не отображается для пользователя</h1>
                )}
                <>
                    {userRole === "admin" && (
                        <>
                        <button className = {visible ? "active" : ""} onClick = {()=>handleChangeVisible(true)}>Включить отображение</button>
                        <button className = {visible ===false   ? "active" : ""} onClick= {()=>handleChangeVisible(false)}>Отключить отображение</button>
                        </>
                        )}
                </>
            </div>
        </div>
    )
}


export default memo(ProductInfoHeader);
