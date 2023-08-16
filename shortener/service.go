package shortener

import "context"

func CreateLink(ctx context.Context, shortURL, longURL string) (Link, error) {
	link, err := NewLink(shortURL, longURL)
	if err != nil {
		return link, err
	}

	tx, err := Pool.Begin(ctx)
	if err != nil {
		return link, err
	}

	err = Save(ctx, tx, link)
	if err != nil {
		tx.Rollback(ctx)
		return link, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return link, err
	}

	return link, nil
}

func GetLink(ctx context.Context, shortURL string) (Link, error) {
	var link Link
	tx, err := Pool.Begin(ctx)
	if err != nil {
		return link, err
	}

	link, err = Find(ctx, tx, shortURL)
	if err != nil {
		tx.Rollback(ctx)
		return link, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return link, err
	}

	return link, nil
}
