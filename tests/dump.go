package main

//func main() {
//	config := config2.NewConfig()
//	logger := lib.NewZapLogger(config)
//	db := infrastructure.NewDatabase(config, logger)
//	count := 0
//	var queue []models.Product
//	ctx := context.WithValue(context.Background(), "x-user-id", "ee564790-1e10-43a0-9968-78dda6496ff9")
//	for true {
//		queue = append(queue, models.Product{
//			Name:        gofakeit.Name(),
//			Description: gofakeit.AdjectiveDescriptive(),
//			Price:       gofakeit.Int64(),
//			Quantity:    gofakeit.Int64(),
//			Images:      []string{gofakeit.ImageURL(100, 100)},
//		})
//		count++
//		wg := sync.WaitGroup{}
//		if len(queue) == 1000 {
//			wg.Add(20)
//			for i := 0; i < 20; i++ {
//				go func() {
//					defer wg.Done()
//					err := db.RDBMS.WithContext(ctx).Create(&queue).Error
//					if err != nil {
//						logger.Error("Error", zap.Error(err))
//					}
//					queue = []models.Product{}
//				}()
//			}
//			wg.Wait()
//		}
//		if count == 100000 {
//			break
//		}
//	}
//}
