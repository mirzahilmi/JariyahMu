package repository

var (
	// Users
	queryCreateUser = `
		INSERT INTO Users (ID, FullName, Email, HashedPassword, ProfilePicture)
		VALUE (:ID, :FullName, :Email, :HashedPassword, :ProfilePicture);
	`
	queryGetUserByParam = `
		SELECT ID, FullName, Email, HashedPassword, ProfilePicture, Active
		FROM Users
		WHERE %s = ?
		LIMIT 1;
	`
	queryUpdateUserPassword = `
		UPDATE Users
		SET HashedPassword = :HashedPassword
		WHERE ID = :ID;
	`
	queryUpdateUserStatus = `
		UPDATE Users 
		SET Active = TRUE
		WHERE ID = ?;
	`
	// UserVerifications
	queryCreateUserVerification = `
		INSERT INTO UserVerifications (ID, UserID, Token)
		VALUE (:ID, :UserID, :Token);
	`
	queryGetUserVerificationByIDAndToken = `
		SELECT ID, UserID, Token
		FROM UserVerifications
		WHERE ID = ? AND Token = ?
		LIMIT 1;
	`
	queryUpdateUserVerificationStatus = `
		UPDATE UserVerifications
		SET Succeed = TRUE
		WHERE ID = ?;
	`
	// ResetAttempts
	queryCreateResetAttempt = `
		INSERT INTO Users (ID, UserID, Token, ValidUntil) 
		VALUE (:ID, :UserID, :Token, :ValidUntil);
	`
	queryDeleteOldResetAttempt = `
		DELETE FROM ResetAttempts
		WHERE Succeed = FALSE AND UserID = (
			SELECT ID FROM Users
			WHERE Email = ?
		);
	`
	queryGetResetAttemptExpiration = `
		SELECT Expiration FROM ResetAttempts
		WHERE UserID = ? AND Token = ?
		LIMIT 1;
	`
	queryUpdateResetAttemptStatus = `
		UPDATE ResetAttempts
		SET Succeed = TRUE
		WHERE ID = ?;
	`
)
