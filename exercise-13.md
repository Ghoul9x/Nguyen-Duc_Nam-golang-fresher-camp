- Nếu chúng ta có nhiều hơn 1 module, làm sao để giao tiếp với nhau. Giả sử module "Restaurant" cần data số lượt like từ module "Like Restaurant" thì sẽ truy xuất như thế nào?

    + Nếu chúng ta có nhiều hơn 1 module, các module ấy muốn giao tiếp với nhau thì phải giao tiếp qua API
    + Module "Restaurant" cần data số lượt like từ module "Like Restaurant" thì chúng ta cần tạo thêm hàm GetRestaurantLikes ở tầng storage trong module "Like Restaurant" để đếm số like, sau đó gọi lại hàm ở tầng biz của module "Restaurant" để lấy dữ liệu từ module "Like Restaurant"