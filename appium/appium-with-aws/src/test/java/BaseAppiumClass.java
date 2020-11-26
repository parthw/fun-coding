
import io.appium.java_client.android.AndroidDriver;
import io.appium.java_client.remote.MobileCapabilityType;
import org.openqa.selenium.remote.DesiredCapabilities;
import org.testng.annotations.AfterTest;
import org.testng.annotations.BeforeTest;

import java.net.URL;
import java.util.concurrent.TimeUnit;

public class BaseAppiumClass {

    static AndroidDriver androidDriver;

    @BeforeTest
    public void setup() {
        try {
            System.out.println("Starting android driver desired capabilities setup");
            DesiredCapabilities capabilities = new DesiredCapabilities();
            capabilities.setCapability(MobileCapabilityType.NEW_COMMAND_TIMEOUT, "60");
            capabilities.setCapability("appPackage", "com.myapp");
            capabilities.setCapability("appActivity",
                    "com.myapp.MainActivity");

            URL url = new URL("http://127.0.0.1:4723/wd/hub");
            androidDriver = new AndroidDriver(url, capabilities);

            System.out.println("Successfully completed android driver desired capabilities setup");

            androidDriver.manage().timeouts().implicitlyWait(10, TimeUnit.SECONDS);

        } catch (Exception exception) {
            System.out.println("Caused by - "+ exception.getCause());
            System.out.println("Message - "+ exception.getMessage());
            exception.printStackTrace();
        }
    }

    @AfterTest
    public void teardown() {
        androidDriver.quit();
    }
}
