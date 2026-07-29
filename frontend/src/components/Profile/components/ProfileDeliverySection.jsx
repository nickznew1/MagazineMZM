import {memo, useEffect, useState} from "react";
import "./style.css"

import {useForm} from "react-hook-form";
import {config} from "../../../config";


const ProfileDelivery =({profile}) => {
    const {
        register,
        handleSubmit,
        formState: { errors },
        reset,
    } = useForm({
        mode: "onBlur",
        values:{
            phone_number:"",
            city:"",
            address:"",
        }
    });



    const { user, setProfile} = profile;

    const [visibleDelivery, setVisibleDelivery] = useState(false)
    const toggleDelivery = () => {
        setVisibleDelivery(!visibleDelivery);
    }

    useEffect(() => {
        if (user?.delivery) {
           reset ({
               phone_number: user.delivery.phone_number || "",
               city: user.delivery.city || "",
               second_name: user.delivery.second_name || "",
           })
        }
    }, [user,reset]);


    const handleSubmitDelivery = async (data) => {
        const payload = {
            id:user.profileData.id,
            ...data
        };
        try {
            const response = await fetch(`${config.apiUrl}/profile/delivery`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(payload),
            })


            setProfile({
                ...user,
                profileData:{
                    ...user.profileData,
                    delivery:payload
                }
            })
            setVisibleDelivery(false);
        } catch (error){
            throw new Error(error)
        }
    }



    const handleUpdateDelivery = async (data) => {
        const payload = {
            id:user.profileData.id,
            ...data,
        };
        try {
            const response = await fetch(`${config.apiUrl}/profile/delivery/up`, {
                method: 'PATCH',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(payload),
            })

            const result = await response.json();
            console.log('Успех:', result);
            setProfile({
                ...user,
                profileData:{
                    ...user.profileData,
                    delivery:payload
                }
            })
            setVisibleDelivery(false);
        } catch (error){
            console.error("error", error);
        }
    }

    return(

        <div className = "user-profile-delivery-wrapper">
        { user.profileData.delivery?.id ? (
            <>
            <div className = "user-profile-change__btn-wrapper">
                    <div className = "user-profile-change__btn">
                    <button type ="button" className ={visibleDelivery ? "active" : ""} onClick = {toggleDelivery}>
                        Изменить данные доставки
                    </button>
                    </div>
            </div>
                    {visibleDelivery ? (
                            <form id ="user-profile-form-change" onSubmit={handleSubmit(handleUpdateDelivery)}>
                                <div className="user-profile-change__inputs-wrapper">
                                <div className="user-profile-change__inputs">
                                    <input type="text"
                                           placeholder ="Введите номер телефона"
                                           style={{ borderColor: errors.phone_number ? 'red' : '#ccc'  }}
                                           {...register("phone_number", {
                                               pattern: { value: /^[0-9]+$/, message: "Только цифры" }
                                           })
                                           }>
                                    </input>
                                    {errors.phone_number && <span className = "user-profile-errors">{errors.phone_number.message}</span> }
                                </div>
                                <div className="user-profile-change__inputs">
                                <input type="text"
                                       placeholder ="Введите город"
                                       style={{ borderColor: errors.city ? 'red' : '#ccc'  }}
                                       {...register("city", {
                                           pattern: { value: /^[а-яА-ЯёЁ]+$/, message: "Город должен быть на русском языке" }
                                       })
                                       }>
                                </input>
                                {errors.city && <span className = "user-profile-errors">{errors.city.message}</span> }
                                </div>
                                <div className="user-profile-change__inputs">
                                <input type="text"
                                       placeholder ="Введите адрес"
                                       style={{ borderColor: errors.address ? 'red' : '#ccc'  }}
                                       {...register("address", {
                                           pattern: { value: /^[а-яА-ЯёЁ]+$/, message: "Адрес должен быть на русском языке" }
                                       })
                                       }>
                                </input>
                                {errors.address && <span className = "user-profile-errors">{errors.address.message}</span> }
                                </div>
                                    <div className ="user-profile-submit__btn-wrapper">
                                    <div className="user-profile-submit__btn">
                                        <button type="submit" form ="user-profile-form-change">Отправить</button>
                                    </div>
                                </div>
                                </div>

                            </form>
                    ) : (
                        <div className = "user-delivery-info">
                            <span>Ваш номер телефона: {user.profileData.delivery.phone_number}</span>
                            <span>Город: {user.profileData.delivery.city}</span>
                            <span>Адрес: {user.profileData.delivery.address}</span>
                        </div>

                    )}
            </>
            ) : (

                <form id ="user-profile-form-create" onSubmit={handleSubmit(handleSubmitDelivery)}>
                    <div className="user-profile-delivery-change__inputs-wrapper">
                    <div className="user-profile-change__inputs">
                        <input type="text"
                               placeholder ="Введите номер"
                               style={{ borderColor: errors.phone_number ? 'red' : '#ccc'  }}
                               {...register("phone_number", {
                                   pattern: { value: /^[0-9]+$/, message: "Только цифры" }
                               })
                               }>
                        </input>
                        {errors.phone_number && <span className = "user-profile-errors">{errors.phone_number.message}</span> }
                    </div>
            <div className="user-profile-change__inputs">
                    <input type="text"
                           placeholder ="Введите город"
                           style={{ borderColor: errors.city ? 'red' : '#ccc'  }}
                           {...register("city", {
                               pattern: { value: /^[а-яА-ЯёЁ]+$/, message: "Город должен быть на русском языке" }
                           })
                           }>
                    </input>
                    {errors.city && <span className = "user-profile-errors">{errors.city.message}</span> }
                    </div>
                    <div className="user-profile-change__inputs">
                    <input type="text"
                           placeholder ="Введите адрес"
                           style={{ borderColor: errors.address ? 'red' : '#ccc'  }}
                           {...register("address", {
                               pattern: { value: /^[а-яА-ЯёЁ]+$/, message: "Адрес должен быть на русском языке" }
                           })
                           }>
                    </input>
                    {errors.address && <span className = "user-profile-errors">{errors.address.message}</span> }
                    </div>
                        <div className="user-profile-submit__btn">
                            <button type="submit">Сохранить данные</button>
                        </div>
                    </div>
                </form>

            )}

        </div>
)}

export default memo(ProfileDelivery);