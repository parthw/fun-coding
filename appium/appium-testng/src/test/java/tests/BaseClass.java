package tests;

import io.appium.java_client.AppiumDriver;
import io.appium.java_client.remote.MobileCapabilityType;
import org.openqa.selenium.remote.DesiredCapabilities;
import org.testng.annotations.AfterTest;
import org.testng.annotations.BeforeTest;
import org.testng.annotations.Test;

import java.net.URL;

public class BaseClass {

    static AppiumDriver appiumDriver;

    @BeforeTest
    public void setup() {
        try {
            DesiredCapabilities capabilities = new DesiredCapabilities();
            capabilities.setCapability(MobileCapabilityType.DEVICE_NAME, "XX");
            capabilities.setCapability(MobileCapabilityType.UDID, "XX");
            capabilities.setCapability(MobileCapabilityType.PLATFORM_NAME, "Android");
            capabilities.setCapability(MobileCapabilityType.PLATFORM_VERSION, "Android 10");
            capabilities.setCapability(MobileCapabilityType.NEW_COMMAND_TIMEOUT, "60");
            capabilities.setCapability("appPackage", "com.android.chrome");
            capabilities.setCapability("appActivity",
                    "org.chromium.chrome.browser.document.ChromeLauncherActivity");

            URL url = new URL("http://127.0.0.1:4723/wd/hub");
            appiumDriver = new AppiumDriver(url, capabilities);

        } catch (Exception exception) {
            System.out.println("Caused by - "+ exception.getCause());
            System.out.println("Message - "+ exception.getMessage());
            exception.printStackTrace();
        }
    }
    @Test
    public void sampleTest() {
        System.out.println("I am running this test");
    }

    @AfterTest
    public void teardown() {
        appiumDriver.close();
        appiumDriver.quit();

    }
}
