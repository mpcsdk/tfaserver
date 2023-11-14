package tfa

// func (s *sTFA) TfaInfo(ctx context.Context, userId string) (*entity.Tfa, error) {
// 	if userId == "" {
// 		return nil, mpccode.ErrArg
// 	}
// 	info, err := service.DB().FetchTfaInfo(ctx, userId)
// 	if err != nil {
// 		err = gerror.Wrap(err, mpccode.ErrDetails(
// 			mpccode.ErrDetail("userId", userId),
// 		))
// 		return nil, err
// 	}
// 	if info == nil {
// 		return nil, nil
// 	}
// 	///
// 	if info.Mail == "" && info.Phone == "" {
// 		return nil, nil
// 	}
// 	///
// 	return info, nil
// }
