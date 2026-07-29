import {memo, useEffect} from "react";
import {useForm} from "react-hook-form";
import {useNavigate} from "react-router";
import "./styles.css"
import {config} from "../../../config";

const CheckoutFields = ({checkout}) =>{
    const {inputs, token} = checkout
    const navigate = useNavigate();

    const {
        register,
        handleSubmit,
        formState: { errors },
        reset,
    } = useForm({
        mode: "onBlur",
        values:{
            company:"",
            first_name:"",
            second_name:"",
            city:"",
            phone_number:"",
            email:"",
            address:"",
            payment_info:"",
        }

    });

    const validateEmail = (value) => {
        if (!/^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,4}$/i.test(value)){
            return "Некорректный email"
        } else if (value === "") {
            return "Поле обязательно"
        }
    }
    useEffect(() => {
        if (inputs?.profileData.user || inputs?.profileData.personal || inputs?.profileData.delivery) {
            reset({
                email: inputs.profileData.user.email || "",
                company:inputs.profileData.personal.company || "",
                first_name:inputs.profileData.personal.first_name|| "",
                second_name:inputs.profileData.personal.second_name || "",
                phone_number: inputs.profileData.delivery.phone_number || "",
                city: inputs.profileData.delivery.city || "",
                address: inputs.profileData.delivery.address || "",
            })
        }
    }, [inputs, reset]);

    const handleCheckout = async (data) => {

        const payload = {
            user_id:inputs.profileData.user.id,
            login:inputs.profileData.user.login,
            email:data.email,
            company:data.company,
            first_name:data.first_name,
            second_name:data.second_name,
            city:data.city,
            phone_number:data.phone_number,
            address:data.address,
            items: inputs.cartData,
        }
        try {
            const response = await fetch (`${config.apiUrl}/applications`, {
                method: 'PUT',
                headers: token ? { Authorization: `Bearer ${token}` } : {},
                body:JSON.stringify(payload),
            })
            if (response.ok) {
                const inputs = await response.json()
                navigate(`applications/${inputs.application_id}`)
            }

        } catch (error) {
            throw error
        }
    }
    return (
        <>
        <div className = "checkout-sumbit-header">
            <h1>Введите ваши данные</h1>
        </div>
    <form id = "checkout-submit" className = "checkout-submit" onSubmit = {handleSubmit(handleCheckout)}>
        <div className ="checkout-inputs">
            <p>Ваше имя: </p>
            <input type="text"
                   placeholder = "Введите имя"
                   style={{ borderColor: errors.first_name ? 'red' : '#ccc'  }}
                   {...register("first_name", {
                       required : "Поле обязательно для заполнения",
                       pattern: { value: /^[а-яА-ЯёЁ]+$/, message: "Имя должно быть на русском языке" }
                   })
                   }>
            </input>
            {errors.first_name && <span className = "checkout-errors">{errors.first_name.message}</span> }
        </div>
        <div className ="checkout-inputs">
            <p>Ваша фамилия: </p>
            <input type="text"
                   placeholder = "Введите фамилию"
                   style={{ borderColor: errors.second_name ? 'red' : '#ccc'  }}
                   {...register("second_name", {
                       required : "Поле обязательно для заполнения",
                       pattern: { value: /^[а-яА-ЯёЁ]+$/, message: "Фамилия должна быть на русском языке" }
                   })
                   }>
            </input>
            {errors.second_name && <span className = "checkout-errors">{errors.second_name.message}</span> }
        </div>
        <div className ="checkout-inputs">
            <p>Ваш email: </p>
            <input type="email"
                   placeholder = "Введите email"
                   style={{ borderColor: errors.email  ? 'red' : '#ccc'  }}
                   {...register("email", {
                       required : "Поле обязательно для заполнения",
                       validate: validateEmail,
                   })
                   }>
            </input>
            {errors.email && <span className = "checkout-errors">{errors.email.message}</span> }
        </div>
        <div className ="checkout-inputs">
            <p>Ваш номер телефона: </p>
            <input type="text"
                   placeholder = "Введите номер телефона"
                   style={{ borderColor: errors.phone_number ? 'red' : '#ccc'  }}
                   {...register("phone_number", {
                       pattern: { value: /^[0-9]+$/, message: "Только цифры" }
                   })
                   }>
            </input>
            {errors.phone_number && <span className = "checkout-errors">{errors.phone_number.message}</span> }
        </div>
        <div className ="checkout-inputs">
            <p>Название компании: </p>
            <input type="text"
                   placeholder = "Введите название компании"
                   style={{ borderColor: errors.company ? 'red' : '#ccc'  }}
                   {...register("company", {
                       required : "Поле обязательно для заполнения",
                   })
                   }>
            </input>
            {errors.company && <span className = "checkout-errors">{errors.company.message}</span> }
        </div>
        <div className ="checkout-inputs">
            <p>Город: </p>
            <input type="text"
                   placeholder = "Введите ваш город"
                   style={{ borderColor: errors.city ? 'red' : '#ccc'  }}
                   {...register("city", {
                       required : "Поле обязательно для заполнения",
                       pattern: { value: /^[а-яА-ЯёЁ]+$/, message: "Город должен быть на русском языке" }
                   })
                   }>
            </input>
            {errors.city && <span className = "checkout-errors">{errors.city.message}</span> }
        </div>
        <div className ="checkout-inputs">
            <p>Адрес: </p>
            <input type="text"
                   placeholder = "Введите ваш адрес"
                   style={{ borderColor: errors.address ? 'red' : '#ccc'  }}
                   {...register("address", {
                       required : "Поле обязательно для заполнения",
                       pattern: { value: /^[а-яА-ЯёЁ]+$/, message: "Адрес должен быть на русском языке" }
                   })
                   }>
            </input>
            {errors.address && <span className = "checkout-errors">{errors.address.message}</span> }
        </div>
        <div className ="checkout-inputs">
            <p>Расчетный счет предприятия: </p>
            <input type="text"
                   placeholder = "Введите счет"
                   style={{ borderColor: errors.payment_info ? 'red' : '#ccc'  }}
                   {...register("payment_info", {
                   })
                   }>
            </input>
            {errors.payment_info && <span className = "checkout-errors">{errors.payment_info.message}</span> }
        </div>
        <button type = "submit">Оформить заявку на заказ</button>
    </form>
        </>
    )
}

export default memo(CheckoutFields);