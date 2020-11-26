import io.appium.java_client.AppiumDriver;
import io.appium.java_client.MobileElement;
import org.openqa.selenium.By;
import org.openqa.selenium.remote.DesiredCapabilities;

import java.net.URL;

public class CalculatorTests {

    static AppiumDriver driver;

    public static void main(String[] args) {
        try {
            openCalculator();
        } catch (Exception exp) {
            System.out.println(exp.getCause());
            System.out.println(exp.getMessage());
            System.out.println(exp.getStackTrace());

        }

    }
    public static void openCalculator() throws Exception{

        DesiredCapabilities capabilities = new DesiredCapabilities();
        capabilities.setCapability("deviceName", "OnePlus 8");
        capabilities.setCapability("uuid", "192.168.1.201:5588");
        capabilities.setCapability("platformName", "Android");
        capabilities.setCapability("platformVersion", "Android 10");

        capabilities.setCapability("appPackage", "com.oneplus.calculator");
        capabilities.setCapability("appActivity", "com.oneplus.calculator.Calculator");

        URL url = new URL("http://127.0.0.1:4723/wd/hub");

        driver = new AppiumDriver(url, capabilities);
        System.out.println("Application Started...");

        MobileElement two = (MobileElement) driver.findElement(By.id("com.oneplus.calculator:id/digit_2"));
        MobileElement three = (MobileElement) driver.findElement(By.id("com.oneplus.calculator:id/digit_3"));
        MobileElement plus = (MobileElement) driver.findElement(By.id("com.oneplus.calculator:id/op_add"));
        MobileElement equals = (MobileElement) driver.findElement(By.id("com.oneplus.calculator:id/eq"));
        MobileElement result = (MobileElement) driver.findElement(By.id("com.oneplus.calculator:id/result"));

        two.click();
        plus.click();
        three.click();
        equals.click();

        String res = result.getText();
        System.out.println(res);
        System.out.println("Completed Operation...");
    }
}
