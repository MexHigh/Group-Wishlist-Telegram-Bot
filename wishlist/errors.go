package wishlist

import "log"

type GenericWishlistError struct {
	Msg string
	Err error
}

func (gwe *GenericWishlistError) Error() string {
	if gwe.Msg == "" { // try to retrieve the inner error message
		log.Println("Warning: Handled WishlistError with empty Msg attribute")
		innerError := gwe.Unwrap()
		if innerError == nil {
			log.Println("Warning: Handled WishlistError with empty Msg attribute and missing inner error (Err==nil)")
			return "NO ERROR MASSAGE"
		}
		return innerError.Error()
	}
	return gwe.Msg
}

func (gwe *GenericWishlistError) Unwrap() error {
	return gwe.Err
}

type InternalError struct {
	GenericWishlistError
}

type NoDatabaseForChatError struct {
	GenericWishlistError
}

type NoWishesForUserError struct {
	GenericWishlistError
}

type WishIDInvalidError struct {
	GenericWishlistError
}

type WishDoesNotExistError struct {
	GenericWishlistError
}

type WishAlreadyFulfilledError struct {
	GenericWishlistError
}
