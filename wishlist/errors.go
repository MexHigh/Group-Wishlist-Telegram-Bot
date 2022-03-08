package wishlist

type GenericWishlistError struct {
	Msg string
	Err error
}

func (gwe *GenericWishlistError) Error() string {
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
