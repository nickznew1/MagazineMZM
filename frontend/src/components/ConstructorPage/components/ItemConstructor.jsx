import "./style.css"
import {useState, useEffect, memo} from "react";
import {useAuth} from "../../../context/AuthContext";
import {useNavigate} from "react-router";
import {config} from "../../../config";
import NotAuthorized from "../../../features/NotAuthorized/NotAuthorized";

const ItemConstructor = ({constructor}) => {
    const {item} = constructor
    const {user} = useAuth();
    const [submitError, setSubmitError] = useState(false)
    const [count, setCount] = useState('');
    const [activeCategory, setActiveCategory] = useState({});
    const navigate = useNavigate()


    useEffect(() => {
        const defaults = {};

        item.prop_name_array.forEach((prop, index) => {
            const values = item.prop_value_array[index].split(",");

            if (values.length === 1) {
                defaults[prop] = values[0];
            }
        });

        setActiveCategory(defaults);
    }, [item]);

    const activeButton = (key, value) => {
        console.log(key, value)
        setActiveCategory((prev) => ({
            ...prev,
            [key]: value,
        }));
        if (submitError) setSubmitError(false)
    }

    const handleCount = (event) => {
        const value = event.target.value
        setCount(value)
        if (Number(value) <= 0) {
            setSubmitError(true)
        } else {
            setSubmitError(false)
        }
    }

    console.log(item)


    const handleSubmitItem = async (event) => {
        const payload = {
            id:user.id,
            item_id:item.id,
            item_picture: item.item_picture,
            item_secondary_type: item.item_secondary_type,
            item_type: item.item_type,
            name:item.name,
            count: Number(count),
            props : activeCategory,
        }
            try {
            event.preventDefault()
                const allCategoriesFilled = item.prop_name_array.every((prop,_)=> activeCategory[prop] !== undefined
            );
                if (!allCategoriesFilled || count<=0) {
                    setSubmitError(true)
                    return
                }

                const response = await fetch(`${config.apiUrl}/cart/add/`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(payload),
                })
                setSubmitError(false)
                navigate("/cart/")
                } catch (error) {
            throw new Error (error)
            }
   }

    return (
        <>
        {user ? (
        <div className = "constructor-page">

            <div className = "constructor-containers">
                <form id = "constructor-submit" onSubmit = {handleSubmitItem} key = {item.id}>
            {item.prop_name_array.map((prop,index)=> (
                item.prop_value_array[index].split(",").length > 1 &&(
                <div className = {submitError && activeCategory[prop] === undefined ? "constructor-item-error" : "constructor-item"} key={prop}>
                <p>{prop}:</p>
                    <div className = "constructor-item-card" key = {index}>
                        {item.prop_value_array[index].split(",").map((value) => (
                            <button
                                key={`${prop}-${value}`}
                                id = {index}
                                type = "button"
                                onClick={()=>activeButton(prop,value)}
                                className = {activeCategory[prop] === value ? "active" : ""}
                            >{value}</button>
                        ))}
                    </div>
                </div>
                )
            ))}
                    <div className = {submitError ? "constructor-item-error" : "constructor-item"}>
                        <p>Введите количество:</p>
                        <input className = "2" type = "text" value ={count} onChange = {handleCount} />
                    </div>

                </form>
            </div>
            <div className = "constructor-submit__button">
                {submitError && <p>Нужно выбрать все поля.</p>}
                <button type ="submit" form = "constructor-submit" > Добавить в корзину</button>
            </div>


        </div>
        ) : (
            <NotAuthorized />
        )}
        </>

    )



}
export default memo(ItemConstructor);