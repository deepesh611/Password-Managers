
def take_input():
    in1 = input("Enter Username: ")
    pwd1 = input("Enter Password: ")


def verify_master():
    master_password = "hello"
    if input("Enter Master Password: ") == master_password:
        return True
    else:
        return False